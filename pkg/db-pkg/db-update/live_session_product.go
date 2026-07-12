package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateLiveSessionProductPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.LiveSessionProductPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"live_session_products"`)

	if req.LivePrice != nil {
		query = query.Set("live_price", *req.LivePrice)
	}
	if req.CfCodeOverride != nil {
		query = query.Set("cf_code_override", *req.CfCodeOverride)
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
		return fmt.Errorf("live session product with id %d not found", id)
	}
	return nil
}
