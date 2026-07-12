package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// MarkWebhookEventProcessed set processed=true — ໃຊ້ຫຼັງລະບົບ (ຫຼື staff) ຈັດການ event ນີ້ແລ້ວ
func MarkWebhookEventProcessed(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"webhook_events"`).Set("processed", true).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("webhook event with id %d not found", id)
	}
	return nil
}
