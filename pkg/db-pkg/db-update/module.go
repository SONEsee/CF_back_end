package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateModulePut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ModuleRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"modules"`).
		Set("module_name", req.ModuleName).
		Set("display_order", req.DisplayOrder).
		Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("module with id %d not found", id)
	}
	return nil
}

func UpdateModulePatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ModulePatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"modules"`)

	if req.ModuleName != nil {
		query = query.Set("module_name", *req.ModuleName)
	}
	if req.DisplayOrder != nil {
		query = query.Set("display_order", *req.DisplayOrder)
	}
	query = query.Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("module with id %d not found", id)
	}
	return nil
}
