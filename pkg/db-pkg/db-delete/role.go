package dbdelete

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func DeleteRole(ctx context.Context, tx dbpkg.DBTX, RoleID int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"Role"`).Set("deleted_at", time.Now()).Where("id=?", RoleID).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed convert for sql %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed Execue Role %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed whit %w id notfound", err)
	}
	return nil
}
