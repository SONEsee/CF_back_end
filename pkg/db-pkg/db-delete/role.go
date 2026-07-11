package dbdelete

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// DeleteRole ລົບແຖວແທ້ (roles ບໍ່ມີ deleted_at) — ຈະ cascade ລົບ permissions ທີ່ຜູກກັບ role ນີ້ນຳ, ແລະ ຈະລົ້ມເຫຼວຖ້າຍັງມີ users ອ້າງອີງ role ນີ້ຢູ່ (users.role_id ບໍ່ມີ ON DELETE CASCADE)
func DeleteRole(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Delete(`"roles"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("role with id %d not found", id)
	}
	return nil
}
