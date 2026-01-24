package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateRoleDetail(ctx context.Context, tx dbpkg.DBTX, req requestbody.RoleDetail) error {
	psql := db.GetPSQLCommand()
	editStr := boolToString(req.Edit)
	detailStr := boolToString(req.Detail)
	query := psql.Insert(`"RoleDetail"`).Columns("sale", "new", "edit", "delete", "detail", "submenu_id", "role_id").Values(req.Sale, req.New, editStr, req.Delele, detailStr, req.SubMenuID, req.RoleID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
