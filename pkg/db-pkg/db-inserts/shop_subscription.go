package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateShopSubscription(ctx context.Context, tx dbpkg.DBTX, req requestbody.ShopSubscriptionRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"shop_subscriptions"`).
		Columns("shop_id", "plan_id", "start_date", "end_date")
	if req.EndDate == "" {
		query = query.Values(req.ShopID, req.PlanID, req.StartDate, nil)
	} else {
		query = query.Values(req.ShopID, req.PlanID, req.StartDate, req.EndDate)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
