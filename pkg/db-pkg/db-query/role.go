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

func GetRoleDBquery(ctx context.Context, id *int, paginatoinParams *pagination.PaginationParams) ([]dbschema.Role, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()
	if id != nil {
		query := psql.Select("id", "role_name", "detail").From(`"Role"`).Where("id=?", *id).Where("deleted_at IS NULL")
		spl, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("fail convert for sql %w", err)
		}
		var item dbschema.Role
		err = dbpkg.DB.QueryRow(ctx, spl, args...).Scan(

			&item.ID,
			&item.RoleName,
			&item.Detail,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("role with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("fail to scan %w", err)
		}

		return []dbschema.Role{item}, nil, nil
	}
	var totalItem int
	QueryCount := psql.Select("COUNT(*)").From(`"Role"`).Where("deleted_at IS NULL ")

	CountSQL, ContARGS, err := QueryCount.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("fail convert count %w", err)
	}
	err = dbpkg.DB.QueryRow(ctx, CountSQL, ContARGS...).Scan(&totalItem)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	query := psql.Select("id", "role_name", "detail").From(`"Role"`).Where("deleted_at IS NULL").OrderBy("id ASC")
	var paginationResult *pagination.PaginationResult
	if paginatoinParams != nil && paginatoinParams.IsValid() {
		query = query.Limit(uint64(paginatoinParams.GetLimit())).Offset(uint64(paginatoinParams.GetOffset()))
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("fail convert sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query taxes: %w", err)
	}
	defer rows.Close()
	var items []dbschema.Role
	for rows.Next() {
		var item dbschema.Role
		err = rows.Scan(
			&item.ID,
			&item.RoleName,
			&item.Detail,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("fail scan for dat %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating tax rows: %w", err)
	}
	if paginatoinParams != nil && paginatoinParams.IsValid() {
		paginationResult = paginatoinParams.CalculatePagination(totalItem, len(items))
	}
	return items, paginationResult, nil

}
