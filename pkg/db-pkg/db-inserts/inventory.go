package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateInventoryForVariant ສ້າງແຖວ inventories ເລີ່ມຕົ້ນ (qty=0 ໝົດ) ໃຫ້ variant ໃໝ່ — ເອີ້ນຈາກ CreateProductVariant ໃນ transaction ດຽວກັນ
func CreateInventoryForVariant(ctx context.Context, tx dbpkg.DBTX, productVariantID int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"inventories"`).
		Columns("product_variant_id", "actual_qty", "reserved_qty", "available_qty").
		Values(productVariantID, 0, 0, 0)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
