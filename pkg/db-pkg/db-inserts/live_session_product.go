package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateLiveSessionProduct(ctx context.Context, tx dbpkg.DBTX, req requestbody.LiveSessionProductRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"live_session_products"`).
		Columns("live_session_id", "product_variant_id", "live_price", "cf_code_override").
		Values(req.LiveSessionID, req.ProductVariantID, req.LivePrice, nullableStr(req.CfCodeOverride))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
