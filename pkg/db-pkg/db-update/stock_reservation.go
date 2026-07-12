package dbupdate

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// UpdateStockReservationStatus ປ່ຽນສະຖານະ (COMPLETED/EXPIRED) — ເອີ້ນຫຼັງອັບເດດ inventories ແລ້ວໃນ transaction ດຽວກັນ
func UpdateStockReservationStatus(ctx context.Context, tx dbpkg.DBTX, id int64, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"stock_reservations"`).
		Set("status", squirrel.Expr("?::reservation_status_enum", status)).
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
		return fmt.Errorf("stock reservation with id %d not found", id)
	}
	return nil
}
