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

func GetMainMenuByID(ctx context.Context, id int) (*dbschema.MainMenuDGSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_menu", "icon_menu").
		From(`"MainMenu"`).
		Where("id = ?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var item dbschema.MainMenuDGSchema
	err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
		&item.ID,
		&item.NameMenu,
		&item.IconMenu,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("main menu with id %d not found", id)
		}
		return nil, err
	}

	return &item, nil
}

func GetAllMainMenus(ctx context.Context) ([]dbschema.MainMenuDGSchema, error) {
	var result []dbschema.MainMenuDGSchema

	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_menu", "icon_menu").
		From(`"MainMenu"`).
		OrderBy("id ASC")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {

		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item dbschema.MainMenuDGSchema
		err := rows.Scan(
			&item.ID,
			&item.NameMenu,
			&item.IconMenu,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return result, nil
}

func GetTestMainmenu(ctx context.Context) ([]dbschema.MainMenuDGSchema, error) {
	var result []dbschema.MainMenuDGSchema
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_menu", "icon_menu").From(`"MainMenu"`).OrderBy("id ASC")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("fail convert to sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("fail excue %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item dbschema.MainMenuDGSchema
		err = rows.Scan(
			&item.ID,
			&item.NameMenu,
			&item.IconMenu,
		)
		if err != nil {
			return nil, fmt.Errorf("fail to scan rows %w", err)
		}
		result = append(result, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fail error %w ", err)
	}
	return result, nil

}

func GetMainMenuDataQuery(ctx context.Context, id *int, paginatoinParams *pagination.PaginationParams) ([]dbschema.MainMenuDGSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()
	if id != nil {
		query := psql.Select("id", "name_menu", "icon_menu").From(`"MainMenu"`).Where("id=?", *id).Where("deleted_at IS NULL")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed convert for sql %w", err)
		}
		var item dbschema.MainMenuDGSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID,
			&item.NameMenu,
			&item.IconMenu,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("fail not data for id %d", *id)
			}
			return nil, nil, fmt.Errorf("failed excue %w", err)
		}
		return []dbschema.MainMenuDGSchema{item}, nil, nil

	}
	var totalItem int
	queryCount := psql.Select("COUNT(*)").From(`"MainMenu"`).Where("deleted_at IS NULL")
	sqlCount, argsCount, err := queryCount.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed convert for sql %w", err)
	}
	err = dbpkg.DB.QueryRow(ctx, sqlCount, argsCount...).Scan(&totalItem)
	if err != nil {
		return nil, nil, fmt.Errorf("failed for excue %w", err)
	}
	query := psql.Select("id", "name_menu", "icon_menu").From(`"MainMenu"`).Where("deleted_at is NULL").OrderBy("id ASC")
	var paginationResult *pagination.PaginationResult
	if paginatoinParams != nil && paginatoinParams.IsValid() {
		query = query.Limit(uint64(paginatoinParams.GetLimit())).Offset(uint64(paginatoinParams.GetOffset()))
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed  convert for sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query submenus: %w", err)
	}
	defer rows.Close()
	var items []dbschema.MainMenuDGSchema
	for rows.Next() {
		var item dbschema.MainMenuDGSchema
		err = rows.Scan(
			&item.ID,
			&item.NameMenu,
			&item.IconMenu,
		)
		if err != nil {

			return nil, nil, fmt.Errorf("failed excue %w", err)
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
