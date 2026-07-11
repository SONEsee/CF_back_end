package dbdelete

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeleteSubMenu ລົບແຖວແທ້ (sub_menus ບໍ່ມີ deleted_at) — ຈະ cascade ລົບ permissions ທີ່ຜູກກັບເມນູຍ່ອຍນີ້ນຳ
func DeleteSubMenu(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Delete(`"sub_menus"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("sub menu with id %d not found", id)
	}
	return nil
}
