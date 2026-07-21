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


func GetSubMenuDataQuery(ctx context.Context, id *int, q *string, paginationParams *pagination.PaginationParams) ([]dbschema.SubMenuDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(
			"s.id", 
			"s.main_menu_id", 
			"COALESCE(m.menu_name, '') AS main_menu_name", // 🟢 LEFT JOIN ດຶງຊື່
			"s.submenu_name", 
			"s.route_path",
		).
			From(`"sub_menus" s`).
			LeftJoin(`"main_menus" m ON s.main_menu_id = m.id`).
			Where(squirrel.Eq{"s.id": *id})

		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}

		var item dbschema.SubMenuDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID, 
			&item.MainMenuID, 
			&item.MainMenuName, // 🟢 Scan ໃສ່ Struct
			&item.SubmenuName, 
			&item.RoutePath,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("sub menu with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.SubMenuDBSchema{item}, nil, nil
	}

	baseCountQuery := psql.Select("COUNT(*)").From(`"sub_menus" s`)
	baseSelectQuery := psql.Select(
		"s.id", 
		"s.main_menu_id", 
		"COALESCE(m.menu_name, '') AS main_menu_name", // 🟢 LEFT JOIN ດຶງຊື່
		"s.submenu_name", 
		"s.route_path",
	).
		From(`"sub_menus" s`).
		LeftJoin(`"main_menus" m ON s.main_menu_id = m.id`)

	// 🔍 ຖ້າມີ Search Query (q)
	if q != nil && *q != "" {
		searchPattern := "%" + *q + "%"
		searchCondition := squirrel.Or{
			squirrel.ILike{"s.submenu_name": searchPattern},
			squirrel.ILike{"s.route_path": searchPattern},
			squirrel.ILike{"m.menu_name": searchPattern}, // 🟢 ຄົ້ນຫາຕາມຊື່ Main Menu ໄດ້ນຳ
		}
		baseCountQuery = baseCountQuery.LeftJoin(`"main_menus" m ON s.main_menu_id = m.id`).Where(searchCondition)
		baseSelectQuery = baseSelectQuery.Where(searchCondition)
	}

	// 2.1 Count Total Rows
	var totalItem int
	countSQL, countArgs, err := baseCountQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	// 2.2 Add Order & Pagination
	query := baseSelectQuery.OrderBy("s.id ASC")
	if paginationParams != nil && paginationParams.IsValid() {
		query = query.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query sub menus: %w", err)
	}
	defer rows.Close()

	var items []dbschema.SubMenuDBSchema
	for rows.Next() {
		var item dbschema.SubMenuDBSchema
		if err := rows.Scan(
			&item.ID, 
			&item.MainMenuID, 
			&item.MainMenuName, // 🟢 Scan ໃສ່ Struct
			&item.SubmenuName, 
			&item.RoutePath,
		); err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %w", err)
	}

	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		paginationResult = paginationParams.CalculatePagination(totalItem, len(items))
	}

	return items, paginationResult, nil
}