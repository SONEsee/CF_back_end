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

var productColumns = []string{"id", "shop_id", "category_id", "product_name", "description", "image_main_url", "is_active", "created_at", "updated_at"}

func scanProduct(row pgx.Row, item *dbschema.ProductDBSchema) error {
	return row.Scan(
		&item.ID, &item.ShopID, &item.CategoryID, &item.ProductName, &item.Description,
		&item.ImageMainURL, &item.IsActive, &item.CreatedAt, &item.UpdatedAt,
	)
}

func GetProductDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.ProductDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(productColumns...).From(`"products"`).Where("id=?", *id).Where("deleted_at IS NULL")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ProductDBSchema
		if err := scanProduct(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("product with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ProductDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"products"`).Where("deleted_at IS NULL").ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(productColumns...).From(`"products"`).Where("deleted_at IS NULL").OrderBy("id ASC")

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
		return nil, nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ProductDBSchema
	for rows.Next() {
		var item dbschema.ProductDBSchema
		if err := scanProduct(rows, &item); err != nil {
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
