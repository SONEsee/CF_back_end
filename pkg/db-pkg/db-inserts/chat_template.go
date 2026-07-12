package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateChatTemplate(ctx context.Context, tx dbpkg.DBTX, req requestbody.ChatTemplateRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"chat_templates"`).
		Columns("shop_id", "trigger_keyword", "response_body").
		Values(req.ShopID, nullableStr(req.TriggerKeyword), req.ResponseBody)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
