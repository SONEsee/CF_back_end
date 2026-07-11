package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateSubscriptionPlanPut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SubscriptionPlanRequestBody) error {
	psql := db.GetPSQLCommand()
	features := req.Features
	if features == "" {
		features = "{}"
	}
	query := psql.Update(`"subscription_plans"`).
		Set("plan_name", req.PlanName).
		Set("price_monthly", req.PriceMonthly).
		Set("max_users", req.MaxUsers).
		Set("max_products", req.MaxProducts).
		Set("features", features).
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
		return fmt.Errorf("subscription plan with id %d not found", id)
	}
	return nil
}

func UpdateSubscriptionPlanPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.SubscriptionPlanPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"subscription_plans"`)

	if req.PlanName != nil {
		query = query.Set("plan_name", *req.PlanName)
	}
	if req.PriceMonthly != nil {
		query = query.Set("price_monthly", *req.PriceMonthly)
	}
	if req.MaxUsers != nil {
		query = query.Set("max_users", *req.MaxUsers)
	}
	if req.MaxProducts != nil {
		query = query.Set("max_products", *req.MaxProducts)
	}
	if req.Features != nil {
		query = query.Set("features", *req.Features)
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
		return fmt.Errorf("subscription plan with id %d not found", id)
	}
	return nil
}
