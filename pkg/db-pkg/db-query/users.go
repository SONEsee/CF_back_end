package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
	"github.com/jackc/pgx/v5"
)

func GetUserDataDBQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.GetUserDataDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "name", "full_name", "user_name", "password", "profile_image", "back_list", "role_id").
			From(`"User"`).Where("id=?", id).
			Where("deleted_at IS NULL")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed convert for sql %w", err)
		}
		var item dbschema.GetUserDataDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID,
			&item.Name,
			&item.FullName,
			&item.UserName,
			&item.Password,
			&item.ProfileImg,
			&item.BackList,
			&item.RoleID,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("fail scan data row error %w", err)
		}
		return []dbschema.GetUserDataDBSchema{item}, nil, nil

	}
	var totalItem int
	queryCount := psql.Select("COUNT(*)").From(`"User"`).Where("deleted_at IS NULL")
	sqlCount, argsCount, err := queryCount.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed for convert for sql %w", err)
	}
	err = dbpkg.DB.QueryRow(ctx, sqlCount, argsCount...).Scan(&totalItem)
	if err != nil {
		return nil, nil, err
	}
	query := psql.Select("id", "name", "full_name", "user_name", "password", "profile_image", "back_list", "role_id").From(`"User"`).Where("deleted_at IS NULL").OrderBy("id ASC")
	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		query = query.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := query.ToSql()
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	var items []dbschema.GetUserDataDBSchema
	for rows.Next() {
		var item dbschema.GetUserDataDBSchema
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.FullName,
			&item.UserName,
			&item.Password,
			&item.ProfileImg,
			&item.BackList,
			&item.RoleID,
		)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, item)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed getl data user %w", err)
	}
	return items, paginationResult, nil

}
func GetUserByUsername(ctx context.Context, tx dbpkg.DBTX, username string) (*dbschema.GetUserDataDBSchema, error) {
	psql := db.GetPSQLCommand()

	query := psql.
		Select("id", "name", "full_name", "user_name", "password", "profile_image", "back_list", "role_id"). // ✅ ປ່ຽນເປັນ back_list
		From(`"User"`).
		Where("user_name = ?", username).
		Where("deleted_at IS NULL")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	var user dbschema.GetUserDataDBSchema
	err = tx.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.FullName,
		&user.UserName,
		&user.Password,
		&user.ProfileImg,
		&user.BackList,
		&user.RoleID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// import (
// 	"context"

// 	"github.com/SONEsee/go-echo/config/db"
// 	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
// 	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
// )

// func GetUserDataDBQuery(ctx context.Context) ([]dbschema.GetUserDataDBSchema, error) {
// 	var res = []dbschema.GetUserDataDBSchema{}
// 	psql := db.GetPSQLCommand()
// 	query := psql.
// 		Select("id", "name", "email").
// 		From("users")
// 	sql, args, err := query.ToSql()
// 	if err != nil {
// 		return res, err
// 	}

// 	rows, err := dbpkg.DB.Query(ctx, sql, args...)
// 	if err != nil {
// 		return res, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var item dbschema.GetUserDataDBSchema
// 		if err := rows.Scan(&item.ID, &item.Name, &item.Email); err != nil {
// 			return res, err
// 		}
// 		res = append(res, item)
// 	}

// 	return res, nil
// }
