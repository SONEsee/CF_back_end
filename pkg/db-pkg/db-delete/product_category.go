package dbdelete

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeleteProductCategory ລົບແຖວແທ້ (product_categories ບໍ່ມີ deleted_at) — sub-category ຈະຖືກ set parent_id=NULL ອັດຕະໂນມັດ (ON DELETE SET NULL), ຈະລົ້ມເຫຼວຖ້າຍັງມີ products ອ້າງອີງ category ນີ້ຢູ່ໂດຍກົງບໍ່ໄດ້ (products.category_id ເປັນ ON DELETE SET NULL ຄືກັນ)
func DeleteProductCategory(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Delete(`"product_categories"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product category with id %d not found", id)
	}
	return nil
}
