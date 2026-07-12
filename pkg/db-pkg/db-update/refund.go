package dbupdate

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// UpdateRefundStatus ປ່ຽນ status — caller (service) ຕ້ອງກວດ transition ຖືກຕ້ອງກ່ອນເອີ້ນ
func UpdateRefundStatus(ctx context.Context, tx dbpkg.DBTX, id int64, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"refunds"`).
		Set("status", squirrel.Expr("?::refund_status_enum", status)).
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
		return fmt.Errorf("refund with id %d not found", id)
	}
	return nil
}
