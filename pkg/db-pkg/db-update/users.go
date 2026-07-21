package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateUser(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.UserRequestBodyPacth) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"users"`)

	if req.ShopID != nil {
		query = query.Set("shop_id", *req.ShopID)
	}
	if req.RoleID != nil {
		query = query.Set("role_id", *req.RoleID)
	}
	if req.Username != nil {
		query = query.Set("username", *req.Username)
	}
	if req.Password != "" {
		query = query.Set("password_hash", req.Password)
	}
	if req.FullName != nil {
		query = query.Set("full_name", *req.FullName)
	}
	if req.Email != nil {
		query = query.Set("email", *req.Email)
	}
	if req.Phone != nil {
		query = query.Set("phone", *req.Phone)
	}
	if req.ProfileImage != nil {
		query = query.Set("profile_image", *req.ProfileImage)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}

// DeactivateUser ໃຊ້ແທນການລົບ (users ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateUser(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"users"`).Set("is_active", false).Set("updated_at", time.Now()).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}
