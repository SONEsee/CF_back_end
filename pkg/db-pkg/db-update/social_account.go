package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateSocialAccountPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SocialAccountPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"social_accounts"`)

	if req.AccountName != nil {
		query = query.Set("account_name", *req.AccountName)
	}
	if req.AccessToken != nil {
		query = query.Set("access_token", *req.AccessToken)
	}
	if req.TokenExpiresAt != nil {
		query = query.Set("token_expires_at", *req.TokenExpiresAt)
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
		return fmt.Errorf("social account with id %d not found", id)
	}
	return nil
}

// DeactivateSocialAccount ໃຊ້ແທນການລົບ (social_accounts ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateSocialAccount(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"social_accounts"`).Set("is_active", false).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("social account with id %d not found", id)
	}
	return nil
}
