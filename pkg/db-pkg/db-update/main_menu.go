package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateMainMenuPut(ctx context.Context, id int64, tx dbpkg.DBTX, req requestbody.MainMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"main_menus"`).
		Set("module_id", req.ModuleID).
		Set("menu_name", req.NameMenu).
		Set("icon_class", req.IconMenu).
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
		return fmt.Errorf("main menu with id %d not found", id)
	}
	return nil
}

func UpdateMainMenuPacth(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.MainMenuPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"main_menus"`)
	if req.ModuleID != nil {
		query = query.Set("module_id", *req.ModuleID)
	}
	if req.NameMenu != nil {
		query = query.Set("menu_name", *req.NameMenu)
	}
	if req.IconMenu != nil {
		query = query.Set("icon_class", *req.IconMenu)
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
		return fmt.Errorf("main menu with id %d not found", id)
	}
	return nil
}
