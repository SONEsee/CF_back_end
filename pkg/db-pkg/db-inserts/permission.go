package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreatePermission(ctx context.Context, tx dbpkg.DBTX, req requestbody.PermissionRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"permissions"`).
		Columns("role_id", "submenu_id", "can_view", "can_create", "can_update", "can_delete").
		Values(req.RoleID, req.SubmenuID, req.CanView, req.CanCreate, req.CanUpdate, req.CanDelete)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
