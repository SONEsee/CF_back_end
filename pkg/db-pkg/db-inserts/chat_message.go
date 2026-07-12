package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateChatMessage(ctx context.Context, tx dbpkg.DBTX, conversationID int, senderType, messageType string, messageBody, attachmentURL string) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"chat_messages"`).
		Columns("conversation_id", "sender_type", "message_type", "message_body", "attachment_url").
		Values(
			conversationID,
			squirrel.Expr("?::sender_type_enum", senderType),
			squirrel.Expr("?::message_type_enum", messageType),
			nullableStr(messageBody),
			nullableStr(attachmentURL),
		)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
