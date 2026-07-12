package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateOrderStatusLog ບັນທຶກການປ່ຽນ status — fromStatus nil ໝາຍວ່າແມ່ນຄັ້ງທຳອິດ (ຕອນສ້າງ order)
func CreateOrderStatusLog(ctx context.Context, tx dbpkg.DBTX, orderID int, fromStatus *string, toStatus string, changedByType string, changedByID *int, note string) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"order_status_logs"`).
		Columns("order_id", "from_status", "to_status", "changed_by_type", "changed_by_id", "note").
		Values(orderID, fromStatus, toStatus, squirrel.Expr("?::changed_by_type_enum", changedByType), changedByID, nullableStr(note))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
