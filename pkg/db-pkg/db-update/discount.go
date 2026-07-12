package dbupdate

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateDiscountPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.DiscountPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"discounts"`)

	if req.DiscountValue != nil {
		query = query.Set("discount_value", *req.DiscountValue)
	}
	if req.MinOrder != nil {
		query = query.Set("min_order", *req.MinOrder)
	}
	if req.UsageLimit != nil {
		query = query.Set("usage_limit", *req.UsageLimit)
	}
	if req.StartAt != nil {
		query = query.Set("start_at", *req.StartAt)
	}
	if req.EndAt != nil {
		query = query.Set("end_at", *req.EndAt)
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
		return fmt.Errorf("discount with id %d not found", id)
	}
	return nil
}

// IncrementDiscountUsage +1 used_count — ເອີ້ນຫຼັງ order ໃຊ້ coupon ນີ້ສຳເລັດ, ຢູ່ໃນ transaction ດຽວກັບການສ້າງ order
func IncrementDiscountUsage(ctx context.Context, tx dbpkg.DBTX, id int) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"discounts"`).Set("used_count", squirrel.Expr("used_count + 1")).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

// DeactivateDiscount ໃຊ້ແທນການລົບ (discounts ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateDiscount(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"discounts"`).Set("is_active", false).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("discount with id %d not found", id)
	}
	return nil
}
