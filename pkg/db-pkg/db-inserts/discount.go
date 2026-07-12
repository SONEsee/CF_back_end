package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateDiscount(ctx context.Context, tx dbpkg.DBTX, req requestbody.DiscountRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"discounts"`).
		Columns("shop_id", "code", "discount_type", "discount_value", "min_order", "usage_limit", "start_at", "end_at").
		Values(
			req.ShopID,
			req.Code,
			squirrel.Expr("?::discount_type_enum", req.DiscountType),
			req.DiscountValue,
			req.MinOrder,
			req.UsageLimit,
			req.StartAt,
			req.EndAt,
		)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
