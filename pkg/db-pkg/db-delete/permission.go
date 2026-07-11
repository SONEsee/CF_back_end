package dbdelete

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeletePermission ລົບແຖວແທ້ (permissions ບໍ່ມີ deleted_at)
func DeletePermission(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Delete(`"permissions"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("permission with id %d not found", id)
	}
	return nil
}
