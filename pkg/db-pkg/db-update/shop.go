package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateShopPut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ShopRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shops"`).
		Set("shop_name", req.ShopName).
		Set("owner_user_id", req.OwnerUserID).
		Set("phone", req.Phone).
		Set("timezone", req.Timezone).
		Set("updated_at", time.Now()).
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
		return fmt.Errorf("shop with id %d not found", id)
	}
	return nil
}

func UpdateShopPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ShopPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shops"`)

	if req.ShopName != nil {
		query = query.Set("shop_name", *req.ShopName)
	}
	if req.OwnerUserID != nil {
		query = query.Set("owner_user_id", *req.OwnerUserID)
	}
	if req.Phone != nil {
		query = query.Set("phone", *req.Phone)
	}
	if req.Timezone != nil {
		query = query.Set("timezone", *req.Timezone)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("shop with id %d not found", id)
	}
	return nil
}

// UpdateShopStatus ໃຊ້ແທນການລົບ (shops ບໍ່ມີ deleted_at) — ປ່ຽນສະຖານະ ACTIVE/SUSPENDED/TRIAL
func UpdateShopStatus(ctx context.Context, tx dbpkg.DBTX, id int64, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shops"`).
		Set("status", squirrel.Expr("?::shop_status_enum", status)).
		Set("updated_at", time.Now()).
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
		return fmt.Errorf("shop with id %d not found", id)
	}
	return nil
}
