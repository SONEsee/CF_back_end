package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateModule(ctx context.Context, tx dbpkg.DBTX, req requestbody.ModuleRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"modules"`).
		Columns("module_name", "display_order").
		Values(req.ModuleName, req.DisplayOrder)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
