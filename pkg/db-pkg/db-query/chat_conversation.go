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

var chatConversationColumns = []string{"id", "social_account_id", "customer_id", "assigned_staff_id", "last_message_preview", "unread_count", "status", "updated_at"}

func scanChatConversation(row pgx.Row, item *dbschema.ChatConversationDBSchema) error {
	return row.Scan(
		&item.ID, &item.SocialAccountID, &item.CustomerID, &item.AssignedStaffID,
		&item.LastMessagePreview, &item.UnreadCount, &item.Status, &item.UpdatedAt,
	)
}

func GetChatConversationDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.ChatConversationDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(chatConversationColumns...).From(`"chat_conversations"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ChatConversationDBSchema
		if err := scanChatConversation(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("chat conversation with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ChatConversationDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"chat_conversations"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(chatConversationColumns...).From(`"chat_conversations"`).OrderBy("updated_at DESC")

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
		return nil, nil, fmt.Errorf("failed to query chat conversations: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ChatConversationDBSchema
	for rows.Next() {
		var item dbschema.ChatConversationDBSchema
		if err := scanChatConversation(rows, &item); err != nil {
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

// FindActiveConversationForUpdate ຄົ້ນຫາ conversation ທີ່ status≠CLOSED ຂອງ (social_account_id, customer_id) ຄູ່ນີ້ ພ້ອມລັອກແຖວ —
// ໃຊ້ໃນ find-or-create flow ຕອນສົ່ງ message. ຄືນ nil (ບໍ່ແມ່ນ error) ຖ້າບໍ່ພົບ, ໃຫ້ caller ຮູ້ວ່າຕ້ອງສ້າງໃໝ່
func FindActiveConversationForUpdate(ctx context.Context, tx dbpkg.DBTX, socialAccountID, customerID int) (*dbschema.ChatConversationDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(chatConversationColumns...).From(`"chat_conversations"`).
		Where("social_account_id=?", socialAccountID).
		Where("customer_id=?", customerID).
		Where("status<>'CLOSED'").
		OrderBy("updated_at DESC").
		Limit(1).
		Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.ChatConversationDBSchema
	err = scanChatConversation(tx.QueryRow(ctx, sql, args...), &item)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
