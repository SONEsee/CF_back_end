package dbupdate

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateTax(ctx context.Context, tx dbpkg.DBTX, taxID int64, req requestbody.TaxRequestBody) error {
	qspl := db.GetPSQLCommand()
	query := qspl.Update(`"Tax"`).Set("name_tax", req.NameTax).Set("value_tax", req.ValueTax).Set("tax_detail", req.TaxDetail).Where("id=?", taxID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
