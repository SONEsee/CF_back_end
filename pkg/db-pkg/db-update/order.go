package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// UpdateOrderStatus ປ່ຽນ current_status — ເອີ້ນຫຼັງ resolve stock_reservations (ຖ້າຈຳເປັນ) ແລ້ວໃນ transaction ດຽວກັນ
func UpdateOrderStatus(ctx context.Context, tx dbpkg.DBTX, id int64, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"orders"`).
		Set("current_status", squirrel.Expr("?::order_status_enum", status)).
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
		return fmt.Errorf("order with id %d not found", id)
	}
	return nil
}
