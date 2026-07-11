package dbupdate

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateShopSubscriptionPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ShopSubscriptionPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shop_subscriptions"`)

	if req.PlanID != nil {
		query = query.Set("plan_id", *req.PlanID)
	}
	if req.EndDate != nil {
		query = query.Set("end_date", *req.EndDate)
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
		return fmt.Errorf("shop subscription with id %d not found", id)
	}
	return nil
}

// UpdateShopSubscriptionStatus ໃຊ້ແທນການລົບ (shop_subscriptions ບໍ່ມີ deleted_at) — ປ່ຽນສະຖານະ ACTIVE/EXPIRED/CANCELLED
func UpdateShopSubscriptionStatus(ctx context.Context, tx dbpkg.DBTX, id int64, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shop_subscriptions"`).
		Set("status", squirrel.Expr("?::subscription_status_enum", status)).
		Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("shop subscription with id %d not found", id)
	}
	return nil
}
