package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateLiveSession(ctx context.Context, tx dbpkg.DBTX, req requestbody.LiveSessionRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"live_sessions"`).
		Columns("social_account_id", "fb_video_id", "session_title").
		Values(req.SocialAccountID, nullableStr(req.FbVideoID), nullableStr(req.SessionTitle))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
