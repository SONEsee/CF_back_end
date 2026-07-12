package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// MarkCommentRawProcessed set is_processed=true — ເອີ້ນຫຼັງສ້າງ comment_intent ໃນ transaction ດຽວກັນ (Step 4)
func MarkCommentRawProcessed(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"comments_raw"`).Set("is_processed", true).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("comment with id %d not found", id)
	}
	return nil
}
