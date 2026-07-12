package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateProductVariantPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ProductVariantPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"product_variants"`)

	if req.VariantName != nil {
		query = query.Set("variant_name", *req.VariantName)
	}
	if req.SkuCode != nil {
		query = query.Set("sku_code", *req.SkuCode)
	}
	if req.CfCode != nil {
		query = query.Set("cf_code", *req.CfCode)
	}
	if req.Barcode != nil {
		query = query.Set("barcode", *req.Barcode)
	}
	if req.Price != nil {
		query = query.Set("price", *req.Price)
	}
	if req.CostPrice != nil {
		query = query.Set("cost_price", *req.CostPrice)
	}
	if req.WeightGrams != nil {
		query = query.Set("weight_grams", *req.WeightGrams)
	}
	if req.IsActive != nil {
		query = query.Set("is_active", *req.IsActive)
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
		return fmt.Errorf("product variant with id %d not found", id)
	}
	return nil
}

// DeactivateProductVariant ໃຊ້ແທນການລົບ (product_variants ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateProductVariant(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"product_variants"`).Set("is_active", false).Set("updated_at", time.Now()).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product variant with id %d not found", id)
	}
	return nil
}
