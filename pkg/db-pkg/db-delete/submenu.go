package dbdelete

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func DeLeteSubMenu(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"SubMenu"`).Set("deleted_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed convert for sql %w", err)
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed execue for data %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed whit %d id notfound", id)
	}
	return nil
}
