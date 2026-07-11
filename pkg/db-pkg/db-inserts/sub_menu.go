package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateSubMenu(ctx context.Context, tx dbpkg.DBTX, req requestbody.SubMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"sub_menus"`).
		Columns("main_menu_id", "submenu_name", "route_path").
		Values(req.MainMenuID, req.SubmenuName, req.RoutePath)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
