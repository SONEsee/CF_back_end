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
	query := psql.Update(`"User"`)
	if req.FullName != nil {
		query.Set("full_name", req.FullName)
	}
	if req.Name != nil {
		query.Set("name", req.Name)
	}
	if req.UserName != nil {
		query.Set("user_name", req.UserName)
	}
	if req.ProfileImg != nil {
		query.Set("profile_image", req.ProfileImg)
	}
	if req.BlackList != nil {
		query.Set("black_list", req.BlackList)
	}
	if req.RoleID != nil {
		query.Set("role_id", req.RoleID)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")
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
