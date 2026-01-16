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
	query := psql.Update(`"User"`).Where("id=?", id).Where("deleted_at IS NULL")
	if req.FullName != nil {
		query = query.Set("full_name", *req.FullName)
	}
	if req.Name != nil {
		query = query.Set("name", *req.Name)
	}
	if req.Password != "" {
		query = query.Set("password", req.Password)
	}
	if req.UserName != nil {
		query = query.Set("user_name", *req.UserName)
	}
	if req.ProfileImg != nil {
		query = query.Set("profile_image", *req.ProfileImg)
	}
	if req.BlackList != nil {
		query = query.Set("black_list", *req.BlackList)
	}
	if req.RoleID != nil {
		query = query.Set("role_id", *req.RoleID)
	}
	query = query.Set("updated_at", time.Now())
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

// func UpdateUser(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.UserRequestBodyPacth) error {
// 	psql := db.GetPSQLCommand()

// 	// ✅ ເລີ່ມ query ດ້ວຍ WHERE ກ່ອນ
// 	query := psql.Update(`"User"`).
// 		Where("id = ?", id).
// 		Where("deleted_at IS NULL")

// 	// ✅ ຕ້ອງ assign query ຄືນທຸກຄັ້ງ ແລະໃຊ້ *pointer
// 	if req.Name != nil {
// 		query = query.Set("name", *req.Name) // ✅ query = ... ແລະ *req.Name
// 	}
// 	if req.FullName != nil {
// 		query = query.Set("full_name", *req.FullName) // ✅
// 	}
// 	if req.UserName != nil {
// 		query = query.Set("user_name", *req.UserName) // ✅
// 	}
// 	if req.Password != "" {
// 		query = query.Set("password", req.Password) // ✅ password ບໍ່ແມ່ນ pointer
// 	}
// 	if req.ProfileImg != nil {
// 		query = query.Set("profile_image", *req.ProfileImg) // ✅
// 	}
// 	if req.BlackList != nil {
// 		query = query.Set("back_list", *req.BlackList) // ✅ (ກວດສອບຊື່: back_list ຫຼື black_list)
// 	}
// 	if req.RoleID != nil {
// 		query = query.Set("role_id", *req.RoleID) // ✅
// 	}

// 	// ✅ Set updated_at
// 	query = query.Set("updated_at", time.Now())

// 	sql, args, err := query.ToSql()
// 	if err != nil {
// 		return fmt.Errorf("failed to build SQL: %w", err)
// 	}

// 	result, err := tx.Exec(ctx, sql, args...)
// 	if err != nil {
// 		return fmt.Errorf("failed to execute update: %w", err)
// 	}

// 	if result.RowsAffected() == 0 {
// 		return fmt.Errorf("user with id %d not found", id)
// 	}

// 	return nil
// }
