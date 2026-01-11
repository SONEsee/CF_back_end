package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func InsertNewUserTx(ctx context.Context, tx dbpkg.DBTX, req requestbody.UserRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.
		Insert(`"User"`).
		Columns("name", "full_name", "user_name", "password", "profile_image", "back_list", "role_id").
		Values(req.Name, req.FullName, req.UserName, req.Password, req.ProfileImg, req.BlackList, req.RoleID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sql, args...)
	return err
}
