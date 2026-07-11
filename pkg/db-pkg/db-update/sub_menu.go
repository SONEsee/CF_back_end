package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateSubMenuPut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SubMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"sub_menus"`).
		Set("main_menu_id", req.MainMenuID).
		Set("submenu_name", req.SubmenuName).
		Set("route_path", req.RoutePath).
		Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("sub menu with id %d not found", id)
	}
	return nil
}

func UpdateSubMenuPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SubMenuPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"sub_menus"`)

	if req.MainMenuID != nil {
		query = query.Set("main_menu_id", *req.MainMenuID)
	}
	if req.SubmenuName != nil {
		query = query.Set("submenu_name", *req.SubmenuName)
	}
	if req.RoutePath != nil {
		query = query.Set("route_path", *req.RoutePath)
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
		return fmt.Errorf("sub menu with id %d not found", id)
	}
	return nil
}
