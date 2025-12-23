package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateSubMenu(ctx context.Context, tx dbpkg.DBTX, req requestbody.SubMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"SubMenu"`).Columns("name_submenu", "url_submenu", "icon_submenu", "action", "main_menu_id").Values(req.NameSubMenu, req.URLSubMenu, req.IconSubMenu, req.Action, req.MainMenuID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
