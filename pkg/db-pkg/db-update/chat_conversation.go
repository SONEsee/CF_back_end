package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// UpdateConversationAfterMessage ອັບເດດ summary fields ຫຼັງບັນທຶກ message ໃໝ່ — incrementUnread=true ສະເພາະຕອນ sender=CUSTOMER
func UpdateConversationAfterMessage(ctx context.Context, tx dbpkg.DBTX, id int64, preview string, incrementUnread bool) error {
	psql := db.GetPSQLCommand()
	var previewVal interface{}
	if preview != "" {
		previewVal = preview
	}
	query := psql.Update(`"chat_conversations"`).
		Set("last_message_preview", previewVal).
		Set("updated_at", time.Now())
	if incrementUnread {
		query = query.Set("unread_count", squirrel.Expr("unread_count + 1"))
	}
	query = query.Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

// MarkConversationRead set unread_count=0 — ໃຊ້ຕອນ staff ເປີດອ່ານ conversation
func MarkConversationRead(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"chat_conversations"`).Set("unread_count", 0).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat conversation with id %d not found", id)
	}
	return nil
}

func UpdateChatConversationPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ChatConversationPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"chat_conversations"`)

	if req.AssignedStaffID != nil {
		query = query.Set("assigned_staff_id", *req.AssignedStaffID)
	}
	if req.Status != nil {
		query = query.Set("status", squirrel.Expr("?::conversation_status_enum", *req.Status))
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat conversation with id %d not found", id)
	}
	return nil
}
