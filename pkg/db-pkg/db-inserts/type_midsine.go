package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateTypeMidsine(ctx context.Context, tx dbpkg.DBTX, req requestbody.TypeMedicine) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"TypeMidisine"`).Columns("name_type", "detail_type").Values(req.NameType, req.DetailType)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
