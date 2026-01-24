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

func GetDataRoleDetail(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.RoleDetail, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "sale", "new", "edit", "delele", "detail", "submenu_id", "role_id").
			From(`"RoleDetail"`).
			Where("id = ?", *id).
			Where("deleted_at IS NULL")

		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}

		var item dbschema.RoleDetail
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID,
			&item.Sale,
			&item.New,
			&item.Edit,
			&item.Delele,
			&item.Detail,
			&item.SubMenuID,
			&item.RoleID,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("no data found for id %d", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.RoleDetail{item}, nil, nil
	}

	var totalItem int
	countQuery := psql.Select("COUNT(*)").
		From(`"RoleDetail"`).
		Where("deleted_at IS NULL")

	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}

	err = dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select("id", "sale", "new", "edit", "delele", "detail", "submenu_id", "role_id").
		From(`"RoleDetail"`).
		Where("deleted_at IS NULL").
		OrderBy("id DESC")

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
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var items []dbschema.RoleDetail
	for rows.Next() {
		var item dbschema.RoleDetail
		err = rows.Scan(
			&item.ID,
			&item.Sale,
			&item.New,
			&item.Edit,
			&item.Delele,
			&item.Detail,
			&item.SubMenuID,
			&item.RoleID,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return items, paginationResult, nil
}
