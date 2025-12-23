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

func GetSubmenuWhitAll(ctx context.Context) ([]dbschema.SubMenuSchema, error) {
	var result []dbschema.SubMenuSchema
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_submenu", "icon_submenu", "url_submenu", "action", "main_menu_id").From(`"SubMenu"`).OrderBy("id ASC")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("fail convert forn sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)

	if err != nil {
		return nil, fmt.Errorf("fail excue %w ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item dbschema.SubMenuSchema
		err = rows.Scan(
			&item.ID,
			&item.NameSubMenu,
			&item.IconSubMenu,
			&item.URLSubMenu,
			&item.Action,
			&item.MainMenuID,
		)
		if err != nil {
			return nil, fmt.Errorf("fail to scan data %w", err)
		}
		result = append(result, item)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("fail to data %w", err)
	}
	return result, nil

}
func GetSubMenuByParam(ctx context.Context, id *int, paginatoinParams *pagination.PaginationParams) ([]dbschema.SubMenuSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()
	if id != nil {
		query := psql.Select("id", "name_submenu", "icon_submenu", "url_submenu", "action", "main_menu_id").From(`"SubMenu"`).Where("id=?", *id).Where("deleted_at IS NULL")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("fail convert for sql %w", err)
		}
		var item dbschema.SubMenuSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID,
			&item.NameSubMenu,
			&item.IconSubMenu,
			&item.URLSubMenu,
			&item.Action,
			&item.MainMenuID,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("fail not data for id %d", err)
			}
			return nil, nil, fmt.Errorf("failed not data %w", err)
		}
		return []dbschema.SubMenuSchema{item}, nil, nil

	}
	var totalItem int
	CountQuery := psql.Select("COUNT(*)").From(`"SubMenu"`).Where("deleted_at IS NULL")
	CountSQL, CountARGS, err := CountQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed convert for sql %w", err)
	}
	err = dbpkg.DB.QueryRow(ctx, CountSQL, CountARGS...).Scan(&totalItem)
	if err != nil {
		return nil, nil, fmt.Errorf("failed for excue %w", err)
	}
	query := psql.Select("id", "name_submenu", "icon_submenu", "url_submenu", "action", "main_menu_id").From(`"SubMenu"`).Where("deleted_at IS NULL").OrderBy("id ASC")
	var paginationResult *pagination.PaginationResult
	if paginatoinParams != nil && paginatoinParams.IsValid() {
		query = query.Limit(uint64(paginatoinParams.GetLimit())).Offset(uint64(paginatoinParams.GetOffset()))
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed for convert for sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query submenus: %w", err)
	}
	defer rows.Close()
	var items []dbschema.SubMenuSchema
	for rows.Next() {
		var item dbschema.SubMenuSchema
		err = rows.Scan(
			&item.ID,
			&item.NameSubMenu,
			&item.IconSubMenu,
			&item.URLSubMenu,
			&item.Action,
			&item.MainMenuID,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed excue for rows data %w", err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating tax rows: %w", err)
	}
	if paginatoinParams != nil && paginatoinParams.IsValid() {
		paginationResult = paginatoinParams.CalculatePagination(totalItem, len(items))
	}
	return items, paginationResult, nil

}
