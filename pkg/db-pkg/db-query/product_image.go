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

func GetProductImageDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.ProductImageDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "product_id", "image_url", "sort_order").From(`"product_images"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ProductImageDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(&item.ID, &item.ProductID, &item.ImageURL, &item.SortOrder)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("product image with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ProductImageDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"product_images"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select("id", "product_id", "image_url", "sort_order").From(`"product_images"`).OrderBy("sort_order ASC", "id ASC")

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
		return nil, nil, fmt.Errorf("failed to query product images: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ProductImageDBSchema
	for rows.Next() {
		var item dbschema.ProductImageDBSchema
		if err := rows.Scan(&item.ID, &item.ProductID, &item.ImageURL, &item.SortOrder); err != nil {
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
