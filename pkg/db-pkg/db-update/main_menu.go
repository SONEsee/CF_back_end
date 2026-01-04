package dbupdate

import (
	"context"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateMainMenuPut(ctx context.Context, id int64, tx dbpkg.DBTX, req requestbody.MainMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"MainMenu"`).Set("name_menu", req.NameMenu).Set("icon_menu", req.IconMenu).Set("updated_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

func UpdateMainMenuPacth(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.MainMenuPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"MainMenu"`)
	if req.NameMenu != nil {
		query = query.Set("name_menu", req.NameMenu)
	}
	if req.IconMenu != nil {
		query = query.Set("icon_menu", req.IconMenu)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
