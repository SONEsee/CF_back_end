package dbupdate

import (
	"context"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateRolePut(ctx context.Context, tx dbpkg.DBTX, idRole int64, req requestbody.Role) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"Role"`).Set("role_name", req.RoleName).Set("detail", req.Detail).Set("updated_at", time.Now()).Where("id=?", idRole).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

func UpdateRolePacth(ctx context.Context, tx dbpkg.DBTX, idRole int64, req requestbody.RolePatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"Role"`)

	if req.RoleName != nil {
		query = query.Set("role_name", *req.RoleName)
	}
	if req.Detail != nil {
		query = query.Set("detail", *req.Detail)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", idRole).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err

}
