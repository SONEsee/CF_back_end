package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateMainMenu(ctx context.Context, tx dbpkg.DBTX, req requestbody.MainMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"MainMenu"`).Columns("name_menu", "icon_menu").Values(req.NameMenu, req.IconMenu)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

func CreateMainMenuTest(ctx context.Context, tx dbpkg.DBTX, req requestbody.MainMenuRequesBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"MainMenu"`).Columns("name_menu", "icon_menu").Values(req.NameMenu, req.IconMenu)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
