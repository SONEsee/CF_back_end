package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateSubscriptionPlan(ctx context.Context, tx dbpkg.DBTX, req requestbody.SubscriptionPlanRequestBody) error {
	psql := db.GetPSQLCommand()
	features := req.Features
	if features == "" {
		features = "{}"
	}
	query := psql.Insert(`"subscription_plans"`).
		Columns("plan_name", "price_monthly", "max_users", "max_products", "features").
		Values(req.PlanName, req.PriceMonthly, req.MaxUsers, req.MaxProducts, features)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
