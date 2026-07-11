package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/jackc/pgx/v5"
)

// GetShopSettingByShopID ດຶງຄ່າຕັ້ງຂອງຮ້ານ (1 shop = 1 ແຖວ, ບໍ່ຕ້ອງມີ pagination)
func GetShopSettingByShopID(ctx context.Context, shopID int) (*dbschema.ShopSettingDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "shop_id", "currency", "vat_rate", "auto_reply_msg", "business_hours", "created_at", "updated_at").
		From(`"shop_settings"`).Where("shop_id=?", shopID)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.ShopSettingDBSchema
	err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
		&item.ID, &item.ShopID, &item.Currency, &item.VatRate, &item.AutoReplyMsg, &item.BusinessHours, &item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("shop setting for shop_id %d not found", shopID)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
