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

var webhookEventColumns = []string{"id", "social_account_id", "event_type", "raw_payload", "processed", "received_at"}

func scanWebhookEvent(row pgx.Row, item *dbschema.WebhookEventDBSchema) error {
	return row.Scan(&item.ID, &item.SocialAccountID, &item.EventType, &item.RawPayload, &item.Processed, &item.ReceivedAt)
}

// GetWebhookEventDataQuery — id ຫາລາຍການດຽວ, ຫຼືກັ່ນຕອງດ້ວຍ processed (nil = ທັງໝົດ, false = ຍັງບໍ່ໄດ້ປະມວນຜົນ)
func GetWebhookEventDataQuery(ctx context.Context, id *int, processed *bool, paginationParams *pagination.PaginationParams) ([]dbschema.WebhookEventDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(webhookEventColumns...).From(`"webhook_events"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.WebhookEventDBSchema
		if err := scanWebhookEvent(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("webhook event with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.WebhookEventDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"webhook_events"`)
	listQuery := psql.Select(webhookEventColumns...).From(`"webhook_events"`)
	if processed != nil {
		countQuery = countQuery.Where("processed=?", *processed)
		listQuery = listQuery.Where("processed=?", *processed)
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
		return nil, nil, fmt.Errorf("failed to query webhook events: %w", err)
	}
	defer rows.Close()

	var items []dbschema.WebhookEventDBSchema
	for rows.Next() {
		var item dbschema.WebhookEventDBSchema
		if err := scanWebhookEvent(rows, &item); err != nil {
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
