package dbupdate

import (
	"context"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateSubMenuPut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SubMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"SubMenu"`).Set("name_submenu", req.NameSubMenu).Set("url_submenu", req.URLSubMenu).Set("icon_submenu", req.IconSubMenu).Set("action", req.Action).Set("main_menu_id", req.MainMenuID).Where("id=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

func UpdateSubMenuPacth(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SubMenuRequesBodyPact) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"SubMenu"`)
	if req.NameSubMenu != nil {
		query = query.Set("name_submenu", *req.NameSubMenu)
	}
	if req.URLSubMenu != nil {
		query = query.Set("url_submenu", *req.URLSubMenu)
	}
	if req.IconSubMenu != nil {
		query = query.Set("icon_submenu", *req.IconSubMenu)
	}
	if req.Action != nil {
		query = query.Set("action", *req.Action)
	}
	if req.MainMenuID != nil {
		query = query.Set("main_menu_id", *req.MainMenuID)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
