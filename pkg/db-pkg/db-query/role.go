package dbquery

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
	"github.com/jackc/pgx/v5"
)

func GetRoleDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams, q string) ([]dbschema.RoleDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "shop_id", "role_name", "description", "created_at").From(`"roles"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.RoleDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(&item.ID, &item.ShopID, &item.RoleName, &item.Description, &item.CreatedAt)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("role with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.RoleDBSchema{item}, nil, nil
	}

	// ຄົ້ນຫາຈາກ role_name ຫຼື description (ILIKE = case-insensitive ໃນ Postgres)
	applySearch := func(b squirrel.SelectBuilder) squirrel.SelectBuilder {
		if q == "" {
			return b
		}
		like := "%" + q + "%"
		return b.Where(squirrel.Or{
			squirrel.ILike{"role_name": like},
			squirrel.ILike{"description": like},
		})
	}

	var totalItem int
	countSQL, countArgs, err := applySearch(psql.Select("COUNT(*)").From(`"roles"`)).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := applySearch(psql.Select("id", "shop_id", "role_name", "description", "created_at").From(`"roles"`)).OrderBy("id ASC")

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
		return nil, nil, fmt.Errorf("failed to query roles: %w", err)
	}
	defer rows.Close()

	var items []dbschema.RoleDBSchema
	for rows.Next() {
		var item dbschema.RoleDBSchema
		if err := rows.Scan(&item.ID, &item.ShopID, &item.RoleName, &item.Description, &item.CreatedAt); err != nil {
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

// GetRoleOptionsQuery ດຶງ role ທັງໝົດແບບບໍ່ມີ limit/pagination — ໃຊ້ສຳລັບ dropdown/autocomplete
func GetRoleOptionsQuery(ctx context.Context) ([]dbschema.RoleOptionDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "role_name").From(`"roles"`).OrderBy("role_name ASC")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query role options: %w", err)
	}
	defer rows.Close()

	var items []dbschema.RoleOptionDBSchema
	for rows.Next() {
		var item dbschema.RoleOptionDBSchema
		if err := rows.Scan(&item.ID, &item.RoleName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return items, nil
}
