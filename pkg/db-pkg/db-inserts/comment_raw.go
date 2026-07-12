package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateCommentRaw(ctx context.Context, tx dbpkg.DBTX, req requestbody.CommentRawRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"comments_raw"`).
		Columns("live_session_id", "fb_comment_id", "fb_user_id", "comment_message").
		Values(req.LiveSessionID, nullableStr(req.FbCommentID), nullableStr(req.FbUserID), nullableStr(req.CommentMessage))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
