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

var commentRawColumns = []string{"id", "live_session_id", "fb_comment_id", "fb_user_id", "comment_message", "is_processed", "received_at"}

func scanCommentRaw(row pgx.Row, item *dbschema.CommentRawDBSchema) error {
	return row.Scan(&item.ID, &item.LiveSessionID, &item.FbCommentID, &item.FbUserID, &item.CommentMessage, &item.IsProcessed, &item.ReceivedAt)
}

// GetCommentRawDataQuery — id ຫາລາຍການດຽວ, ຫຼືກັ່ນຕອງດ້ວຍ liveSessionID+processed
func GetCommentRawDataQuery(ctx context.Context, id *int64, liveSessionID *int, processed *bool, paginationParams *pagination.PaginationParams) ([]dbschema.CommentRawDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(commentRawColumns...).From(`"comments_raw"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.CommentRawDBSchema
		if err := scanCommentRaw(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("comment with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.CommentRawDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"comments_raw"`)
	listQuery := psql.Select(commentRawColumns...).From(`"comments_raw"`)
	if liveSessionID != nil {
		countQuery = countQuery.Where("live_session_id=?", *liveSessionID)
		listQuery = listQuery.Where("live_session_id=?", *liveSessionID)
	}
	if processed != nil {
		countQuery = countQuery.Where("is_processed=?", *processed)
		listQuery = listQuery.Where("is_processed=?", *processed)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("received_at ASC")
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
		return nil, nil, fmt.Errorf("failed to query comments: %w", err)
	}
	defer rows.Close()

	var items []dbschema.CommentRawDBSchema
	for rows.Next() {
		var item dbschema.CommentRawDBSchema
		if err := scanCommentRaw(rows, &item); err != nil {
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

// GetCommentRawByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນສ້າງ comment_intent ໃນ transaction ດຽວ (Step 4)
func GetCommentRawByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int64) (*dbschema.CommentRawDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(commentRawColumns...).From(`"comments_raw"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.CommentRawDBSchema
	if err := scanCommentRaw(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("comment with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
