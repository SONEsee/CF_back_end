package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateRole(ctx context.Context, tx dbpkg.DBTX, req requestbody.Role) error {
	qspl := db.GetPSQLCommand()
	query := qspl.Insert(`"Role"`).Columns("role_name", "detail").Values(req.RoleName, req.Detail)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
