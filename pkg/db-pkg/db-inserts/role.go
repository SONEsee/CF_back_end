package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateRole(ctx context.Context, tx dbpkg.DBTX, req requestbody.RoleRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"roles"`).
		Columns("shop_id", "role_name", "description").
		Values(req.ShopID, req.RoleName, req.Description)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
