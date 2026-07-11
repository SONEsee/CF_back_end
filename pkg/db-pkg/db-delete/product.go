package dbdelete

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeleteProduct soft-delete ຈິງ (products ມີ column deleted_at)
func DeleteProduct(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"products"`).Set("deleted_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product with id %d not found", id)
	}
	return nil
}
