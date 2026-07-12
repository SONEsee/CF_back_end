package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateWebhookEvent(ctx context.Context, tx dbpkg.DBTX, req requestbody.WebhookEventRequestBody) error {
	psql := db.GetPSQLCommand()
	payload := req.RawPayload
	if payload == "" {
		payload = "{}"
	}
	query := psql.Insert(`"webhook_events"`).
		Columns("social_account_id", "event_type", "raw_payload").
		Values(req.SocialAccountID, nullableStr(req.EventType), payload)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
