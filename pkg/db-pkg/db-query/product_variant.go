package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
	"github.com/jackc/pgx/v5"
)

var productVariantColumns = []string{"id", "product_id", "variant_name", "sku_code", "cf_code", "barcode", "price", "cost_price", "weight_grams", "is_active", "created_at", "updated_at"}

func scanProductVariant(row pgx.Row, item *dbschema.ProductVariantDBSchema) error {
	return row.Scan(
		&item.ID, &item.ProductID, &item.VariantName, &item.SkuCode, &item.CfCode, &item.Barcode,
		&item.Price, &item.CostPrice, &item.WeightGrams, &item.IsActive, &item.CreatedAt, &item.UpdatedAt,
	)
}

func GetProductVariantDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.ProductVariantDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(productVariantColumns...).From(`"product_variants"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ProductVariantDBSchema
		if err := scanProductVariant(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("product variant with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ProductVariantDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"product_variants"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(productVariantColumns...).From(`"product_variants"`).OrderBy("id ASC")

	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		query = query.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query product variants: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ProductVariantDBSchema
	for rows.Next() {
		var item dbschema.ProductVariantDBSchema
		if err := scanProductVariant(rows, &item); err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if paginationParams != nil && paginationParams.IsValid() {
		paginationResult = paginationParams.CalculatePagination(totalItem, len(items))
	}
	return items, paginationResult, nil
}
