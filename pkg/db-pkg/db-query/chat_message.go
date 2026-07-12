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

var chatMessageColumns = []string{"id", "conversation_id", "sender_type", "message_type", "message_body", "attachment_url", "is_read", "sent_at"}

func scanChatMessage(row pgx.Row, item *dbschema.ChatMessageDBSchema) error {
	return row.Scan(
		&item.ID, &item.ConversationID, &item.SenderType, &item.MessageType,
		&item.MessageBody, &item.AttachmentURL, &item.IsRead, &item.SentAt,
	)
}

// GetChatMessageDataQuery — id ຫາລາຍການດຽວ, ຫຼືກັ່ນຕອງດ້ວຍ conversationID ເບິ່ງປະຫວັດແຊັດທັງໝົດ
func GetChatMessageDataQuery(ctx context.Context, id *int64, conversationID *int, paginationParams *pagination.PaginationParams) ([]dbschema.ChatMessageDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(chatMessageColumns...).From(`"chat_messages"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ChatMessageDBSchema
		if err := scanChatMessage(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("chat message with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ChatMessageDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"chat_messages"`)
	listQuery := psql.Select(chatMessageColumns...).From(`"chat_messages"`)
	if conversationID != nil {
		countQuery = countQuery.Where("conversation_id=?", *conversationID)
		listQuery = listQuery.Where("conversation_id=?", *conversationID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("sent_at ASC")
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
		return nil, nil, fmt.Errorf("failed to query chat messages: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ChatMessageDBSchema
	for rows.Next() {
		var item dbschema.ChatMessageDBSchema
		if err := scanChatMessage(rows, &item); err != nil {
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
