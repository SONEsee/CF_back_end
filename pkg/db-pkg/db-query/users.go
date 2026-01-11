package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
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
