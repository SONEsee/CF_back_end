package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateProductVariant ສ້າງ variant ແລ້ວຄືນ id ໃໝ່ (ໃຊ້ RETURNING) — ໃຫ້ service ເອົາ id ນີ້ໄປສ້າງ inventories ຕໍ່ໃນ transaction ດຽວກັນ
func CreateProductVariant(ctx context.Context, tx dbpkg.DBTX, req requestbody.ProductVariantRequestBody) (int64, error) {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"product_variants"`).
		Columns("product_id", "variant_name", "sku_code", "cf_code", "barcode", "price", "cost_price", "weight_grams").
		Values(req.ProductID, req.VariantName, req.SkuCode, req.CfCode, req.Barcode, req.Price, req.CostPrice, req.WeightGrams).
		Suffix("RETURNING id")
	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	if err := tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
