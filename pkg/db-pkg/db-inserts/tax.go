package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateTax(ctx context.Context, tx dbpkg.DBTX, req requestbody.TaxRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"Tax"`).Columns("name_tax", "value_tax", "tax_detail").Values(req.NameTax, req.ValueTax, req.TaxDetail)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
