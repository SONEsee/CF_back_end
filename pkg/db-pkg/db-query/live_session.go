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

var liveSessionColumns = []string{"id", "social_account_id", "fb_video_id", "session_title", "status", "started_at", "ended_at"}

func scanLiveSession(row pgx.Row, item *dbschema.LiveSessionDBSchema) error {
	return row.Scan(&item.ID, &item.SocialAccountID, &item.FbVideoID, &item.SessionTitle, &item.Status, &item.StartedAt, &item.EndedAt)
}

func GetLiveSessionDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.LiveSessionDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(liveSessionColumns...).From(`"live_sessions"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.LiveSessionDBSchema
		if err := scanLiveSession(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("live session with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.LiveSessionDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"live_sessions"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(liveSessionColumns...).From(`"live_sessions"`).OrderBy("id DESC")

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
		return nil, nil, fmt.Errorf("failed to query live sessions: %w", err)
	}
	defer rows.Close()

	var items []dbschema.LiveSessionDBSchema
	for rows.Next() {
		var item dbschema.LiveSessionDBSchema
		if err := scanLiveSession(rows, &item); err != nil {
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
