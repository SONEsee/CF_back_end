package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func InsertNewUserTx(ctx context.Context, tx dbpkg.DBTX, req requestbody.UserRequestBody) error {
	psql := db.GetPSQLCommand()
	var profileImage *string
	if req.ProfileImage != "" {
		profileImage = &req.ProfileImage
	}
	query := psql.Insert(`"users"`).
		Columns("shop_id", "role_id", "username", "password_hash", "full_name", "email", "phone", "profile_image").
		Values(req.ShopID, req.RoleID, req.Username, req.Password, req.FullName, req.Email, req.Phone, profileImage)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
