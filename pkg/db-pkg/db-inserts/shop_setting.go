package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateShopSetting(ctx context.Context, tx dbpkg.DBTX, req requestbody.ShopSettingRequestBody) error {
	psql := db.GetPSQLCommand()
	currency := req.Currency
	if currency == "" {
		currency = "LAK"
	}
	businessHours := req.BusinessHours
	if businessHours == "" {
		businessHours = "{}"
	}
	query := psql.Insert(`"shop_settings"`).
		Columns("shop_id", "currency", "vat_rate", "auto_reply_msg", "business_hours").
		Values(req.ShopID, currency, req.VatRate, req.AutoReplyMsg, businessHours)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
