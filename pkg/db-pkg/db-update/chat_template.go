package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateChatTemplatePatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ChatTemplatePatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"chat_templates"`)

	if req.TriggerKeyword != nil {
		query = query.Set("trigger_keyword", *req.TriggerKeyword)
	}
	if req.ResponseBody != nil {
		query = query.Set("response_body", *req.ResponseBody)
	}
	if req.IsActive != nil {
		query = query.Set("is_active", *req.IsActive)
	}
	query = query.Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("chat template with id %d not found", id)
	}
	return nil
}
