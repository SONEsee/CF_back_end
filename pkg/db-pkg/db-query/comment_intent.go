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

var commentIntentColumns = []string{"id", "comment_raw_id", "customer_id", "matched_product_variant_id", "parsed_qty", "intent_status", "processed_at"}

func scanCommentIntent(row pgx.Row, item *dbschema.CommentIntentDBSchema) error {
	return row.Scan(&item.ID, &item.CommentRawID, &item.CustomerID, &item.MatchedProductVariantID, &item.ParsedQty, &item.IntentStatus, &item.ProcessedAt)
}

func GetCommentIntentDataQuery(ctx context.Context, id *int64, customerID *int, paginationParams *pagination.PaginationParams) ([]dbschema.CommentIntentDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(commentIntentColumns...).From(`"comment_intents"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.CommentIntentDBSchema
		if err := scanCommentIntent(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("comment intent with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.CommentIntentDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"comment_intents"`)
	listQuery := psql.Select(commentIntentColumns...).From(`"comment_intents"`)
	if customerID != nil {
		countQuery = countQuery.Where("customer_id=?", *customerID)
		listQuery = listQuery.Where("customer_id=?", *customerID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("id DESC")
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
		return nil, nil, fmt.Errorf("failed to query comment intents: %w", err)
	}
	defer rows.Close()

	var items []dbschema.CommentIntentDBSchema
	for rows.Next() {
		var item dbschema.CommentIntentDBSchema
		if err := scanCommentIntent(rows, &item); err != nil {
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
