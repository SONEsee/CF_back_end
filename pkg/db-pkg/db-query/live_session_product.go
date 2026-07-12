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

var liveSessionProductColumns = []string{"id", "live_session_id", "product_variant_id", "live_price", "cf_code_override"}

func scanLiveSessionProduct(row pgx.Row, item *dbschema.LiveSessionProductDBSchema) error {
	return row.Scan(&item.ID, &item.LiveSessionID, &item.ProductVariantID, &item.LivePrice, &item.CfCodeOverride)
}

func GetLiveSessionProductDataQuery(ctx context.Context, id *int, liveSessionID *int, paginationParams *pagination.PaginationParams) ([]dbschema.LiveSessionProductDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(liveSessionProductColumns...).From(`"live_session_products"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.LiveSessionProductDBSchema
		if err := scanLiveSessionProduct(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("live session product with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.LiveSessionProductDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"live_session_products"`)
	listQuery := psql.Select(liveSessionProductColumns...).From(`"live_session_products"`)
	if liveSessionID != nil {
		countQuery = countQuery.Where("live_session_id=?", *liveSessionID)
		listQuery = listQuery.Where("live_session_id=?", *liveSessionID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("id ASC")
	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		listQuery = listQuery.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := listQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query live session products: %w", err)
	}
	defer rows.Close()

	var items []dbschema.LiveSessionProductDBSchema
	for rows.Next() {
		var item dbschema.LiveSessionProductDBSchema
		if err := scanLiveSessionProduct(rows, &item); err != nil {
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
