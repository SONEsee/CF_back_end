package dbdelete

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func DeletdedUser(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	qspl := db.GetPSQLCommand()

	query := qspl.Update(`"User"`).Where("id=?", id).Set("deleted_at", time.Now()).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("fail not data for id: %d", err)
	}
	return nil
}
