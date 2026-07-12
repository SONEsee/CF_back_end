package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateSocialAccount(ctx context.Context, tx dbpkg.DBTX, req requestbody.SocialAccountRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"social_accounts"`).
		Columns("shop_id", "platform", "platform_account_id", "account_name", "access_token", "token_expires_at").
		Values(
			req.ShopID,
			squirrel.Expr("?::platform_enum", req.Platform),
			req.PlatformAccountID,
			nullableStr(req.AccountName),
			nullableStr(req.AccessToken),
			req.TokenExpiresAt,
		)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
