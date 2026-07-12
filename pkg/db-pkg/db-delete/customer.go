package dbdelete

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeleteCustomer ລົບແຖວແທ້ (customers ບໍ່ມີ deleted_at) — ຈະ cascade ລົບ customer_addresses ນຳ,
// ແຕ່ຈະລົ້ມເຫຼວຖ້າຍັງມີ orders ອ້າງອີງ customer ນີ້ຢູ່ (orders.customer_id ບໍ່ມີ ON DELETE cascade) — ພຶດຕິກຳນີ້ຕັ້ງໃຈ
func DeleteCustomer(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Delete(`"customers"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("customer with id %d not found", id)
	}
	return nil
}
