package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateTypeMidsinePut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.TypeMedicine) error {
	psql := db.GetPSQLCommand()

	query := psql.Update(`"TypeMidisine"`).
		Set("name_type", req.NameType).
		Set("detail_type", req.DetailType).
		Set("updated_at", "NOW()").
		Where("id_type=?", id).
		Where("deleted_at IS NULL")

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL: %w", err)
	}

	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("type medicine with id %d not found", id)
	}

	return nil
}

func UpdateTypeMidisinePatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.TypeMedisinePatch) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"TypeMidisine"`)
	if req.NameType != nil {
		query = query.Set("name_type", *req.NameType)
	}
	if req.DetailType != nil {
		query = query.Set("detail_type", req.DetailType)
	}
	query = query.Set("updated_at", time.Now()).Where("id_type=?", id).Where("deleted_at IS NULL")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if result.RowsAffected() == 0 {
		return fmt.Errorf("update error for id not found databases %d ", id)
	}

	return nil
}
