package dbdelete

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeleteCustomerAddress ລົບແຖວແທ້ (customer_addresses ບໍ່ມີ deleted_at) — ບໍ່ກວດ is_default ຢູ່ນີ້,
// caller (service) ຕ້ອງກວດ is_default ກ່ອນເອີ້ນ ຢູ່ໃນ transaction ດຽວກັນ ເພື່ອປ້ອງກັນ race condition
func DeleteCustomerAddress(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Delete(`"customer_addresses"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("customer address with id %d not found", id)
	}
	return nil
}
