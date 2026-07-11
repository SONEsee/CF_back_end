package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateRolePut(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.RoleRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"roles"`).
		Set("shop_id", req.ShopID).
		Set("role_name", req.RoleName).
		Set("description", req.Description).
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
		return fmt.Errorf("role with id %d not found", id)
	}
	return nil
}

func UpdateRolePatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.RolePatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"roles"`)

	if req.ShopID != nil {
		query = query.Set("shop_id", *req.ShopID)
	}
	if req.RoleName != nil {
		query = query.Set("role_name", *req.RoleName)
	}
	if req.Description != nil {
		query = query.Set("description", *req.Description)
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
		return fmt.Errorf("role with id %d not found", id)
	}
	return nil
}
