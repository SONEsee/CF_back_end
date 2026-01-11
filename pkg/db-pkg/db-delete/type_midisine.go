package dbdelete

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func DeletedTypeMisine(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"TypeMidisine"`).Set("deleted_at", time.Now()).Where("id_type=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed convert for sql %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed excue %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("error id %d not found", err)
	}
	return nil
}
