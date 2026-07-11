package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdatePermissionPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.PermissionPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"permissions"`)

	if req.CanView != nil {
		query = query.Set("can_view", *req.CanView)
	}
	if req.CanCreate != nil {
		query = query.Set("can_create", *req.CanCreate)
	}
	if req.CanUpdate != nil {
		query = query.Set("can_update", *req.CanUpdate)
	}
	if req.CanDelete != nil {
		query = query.Set("can_delete", *req.CanDelete)
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
		return fmt.Errorf("permission with id %d not found", id)
	}
	return nil
}
