package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateShopSettingPatch(ctx context.Context, tx dbpkg.DBTX, shopID int64, req requestbody.ShopSettingPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shop_settings"`)

	if req.Currency != nil {
		query = query.Set("currency", *req.Currency)
	}
	if req.VatRate != nil {
		query = query.Set("vat_rate", *req.VatRate)
	}
	if req.AutoReplyMsg != nil {
		query = query.Set("auto_reply_msg", *req.AutoReplyMsg)
	}
	if req.BusinessHours != nil {
		query = query.Set("business_hours", *req.BusinessHours)
	}
	query = query.Set("updated_at", time.Now()).Where("shop_id=?", shopID)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("shop setting for shop_id %d not found", shopID)
	}
	return nil
}
