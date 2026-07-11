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

func GetPermissionDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.PermissionDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "role_id", "submenu_id", "can_view", "can_create", "can_update", "can_delete").
			From(`"permissions"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.PermissionDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID, &item.RoleID, &item.SubmenuID, &item.CanView, &item.CanCreate, &item.CanUpdate, &item.CanDelete,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("permission with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.PermissionDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"permissions"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select("id", "role_id", "submenu_id", "can_view", "can_create", "can_update", "can_delete").
		From(`"permissions"`).OrderBy("id ASC")

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
		return nil, nil, fmt.Errorf("failed to query permissions: %w", err)
	}
	defer rows.Close()

	var items []dbschema.PermissionDBSchema
	for rows.Next() {
		var item dbschema.PermissionDBSchema
		if err := rows.Scan(&item.ID, &item.RoleID, &item.SubmenuID, &item.CanView, &item.CanCreate, &item.CanUpdate, &item.CanDelete); err != nil {
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
