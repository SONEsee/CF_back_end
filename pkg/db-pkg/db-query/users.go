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

var userColumns = []string{"id", "shop_id", "role_id", "username", "password_hash", "full_name", "email", "phone", "is_active", "last_login_at", "created_at", "updated_at"}

func scanUser(row pgx.Row, item *dbschema.UserDBSchema) error {
	return row.Scan(
		&item.ID, &item.ShopID, &item.RoleID, &item.Username, &item.PasswordHash,
		&item.FullName, &item.Email, &item.Phone, &item.IsActive, &item.LastLoginAt,
		&item.CreatedAt, &item.UpdatedAt,
	)
}

func GetUserDataDBQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.UserDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(userColumns...).From(`"users"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.UserDBSchema
		if err := scanUser(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("user with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.UserDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"users"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(userColumns...).From(`"users"`).OrderBy("id ASC")

	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		query = query.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var items []dbschema.UserDBSchema
	for rows.Next() {
		var item dbschema.UserDBSchema
		if err := scanUser(rows, &item); err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if paginationParams != nil && paginationParams.IsValid() {
		paginationResult = paginationParams.CalculatePagination(totalItem, len(items))
	}
	return items, paginationResult, nil
}

func GetUserByUsername(ctx context.Context, tx dbpkg.DBTX, username string) (*dbschema.UserDBSchema, error) {

    psql := db.GetPSQLCommand()

    query := psql.
        Select(userColumns...).
        From(`"users"`).
        Where("username = ?", username)

    sql, args, err := query.ToSql()
    if err != nil {
        return nil, err
    }

    fmt.Println("====================")
    fmt.Println("SQL :", sql)
    fmt.Println("ARGS:", args)
    fmt.Println("====================")

    var item dbschema.UserDBSchema

    err = scanUser(tx.QueryRow(ctx, sql, args...), &item)

    fmt.Println("SCAN ERROR =", err)
    fmt.Printf("USER = %+v\n", item)

    if err != nil {
        return nil, err
    }

    return &item, nil
}