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

func GetDataTax(ctx context.Context) ([]dbschema.Tax, error) {
	var result []dbschema.Tax
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_tax", "value_tax", "tax_detail").From(`"Tax"`).OrderBy("id ASC")
	spl, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("fail convert for sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, spl, args...)
	if err != nil {
		return nil, fmt.Errorf("fail Excue data rows %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item dbschema.Tax
		err := rows.Scan(
			&item.ID,
			&item.NameTax,
			&item.ValueTax,
			&item.TaxDetail,
		)
		if err != nil {
			return nil, fmt.Errorf("fail scan data %w ", err)
		}
		result = append(result, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fail not row: %w ", err)
	}
	return result, nil
}

func GetByIdTax(ctx context.Context, id int) (*dbschema.Tax, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_tax", "value_tax", "tax_detail").From(`"Tax"`).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("fail convert for sql %w", err)
	}
	var item dbschema.Tax
	err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
		&item.ID,
		&item.NameTax,
		&item.ValueTax,
		&item.TaxDetail,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("main menu with id %d not found", id)
		}
		return nil, err
	}
	return &item, nil

}

func GetTax(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.Tax, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	// ຖ້າມີ ID ສະເພາະ - ບໍ່ໃຊ້ pagination
	if id != nil {
		query := psql.Select("id", "name_tax", "value_tax", "tax_detail").
			From(`"Tax"`).
			Where("deleted_at IS NULL").
			Where("id = ?", *id)

		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build SQL query: %w", err)
		}

		var item dbschema.Tax
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID,
			&item.NameTax,
			&item.ValueTax,
			&item.TaxDetail,
		)

		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("tax with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to get tax: %w", err)
		}

		return []dbschema.Tax{item}, nil, nil
	}

	// 1. ນັບຈຳນວນທັງໝົດກ່ອນ
	var totalItems int
	countQuery := psql.Select("COUNT(*)").
		From(`"Tax"`).
		Where("deleted_at IS NULL")

	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}

	err = dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItems)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count taxes: %w", err)
	}

	// 2. ສ້າງ query ສຳລັບດຶງຂໍ້ມູນ
	query := psql.Select("id", "name_tax", "value_tax", "tax_detail").
		From(`"Tax"`).
		Where("deleted_at IS NULL").
		OrderBy("name_tax ASC")

	// 3. ເພີ່ມ pagination ຖ້າມີ
	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		// ເພີ່ມ LIMIT ແລະ OFFSET
		query = query.
			Limit(uint64(paginationParams.GetLimit())).
			Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	// 4. ດຶງຂໍ້ມູນ
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query taxes: %w", err)
	}
	defer rows.Close()

	var items []dbschema.Tax
	for rows.Next() {
		var item dbschema.Tax
		err := rows.Scan(
			&item.ID,
			&item.NameTax,
			&item.ValueTax,
			&item.TaxDetail,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan tax row: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating tax rows: %w", err)
	}

	// 5. ຄິດໄລ່ pagination result
	if paginationParams != nil && paginationParams.IsValid() {
		paginationResult = paginationParams.CalculatePagination(totalItems, len(items))
	}

	return items, paginationResult, nil
}

func GetTaxDataByidTest(ctx context.Context, id *int, paginationParam *pagination.PaginationParams) ([]dbschema.Tax, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "name_tax", "value_tax", "tax_detail").From(`"Tax"`).Where("id=?", *id).Where("deleted_at IS NULL")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("fail to convet for sql %w", err)
		}
		var item dbschema.Tax
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID,
			&item.NameTax,
			&item.ValueTax,
			&item.TaxDetail,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("tax with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("fail get tax %w", err)
		}
		return []dbschema.Tax{item}, nil, nil
	}
	var Totalitems int
	countQuery := psql.Select("COUNT(*)").From(`"Tax"`).Where("deleted_at IS NULL").OrderBy("name_tax ASC")
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("fail convert cout totalItem %w ", err)
	}
	// err = dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&Totalitems)
	err = dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&Totalitems)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count taxes: %w ", err)
	}
	query := psql.Select("id", "name_tax", "value_tax", "tax_detail").
		From(`"Tax"`). // ← ແກ້ເປັນ "Tax"
		Where("deleted_at IS NULL").
		OrderBy("name_tax ASC")
	var paginationResult *pagination.PaginationResult
	if paginationParam != nil && paginationParam.IsValid() {

		query = query.
			Limit(uint64(paginationParam.GetLimit())).
			Offset(uint64(paginationParam.GetOffset()))
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("fail convert for sql query %w", err)
	}

	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query taxes: %w", err)
	}
	defer rows.Close()
	var items []dbschema.Tax
	for rows.Next() {
		var item dbschema.Tax
		err := rows.Scan(
			&item.ID,
			&item.NameTax,
			&item.ValueTax,
			&item.TaxDetail,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("fail to scan data %w", err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating tax rows: %w", err)
	}
	if paginationParam != nil && paginationParam.IsValid() {
		paginationResult = paginationParam.CalculatePagination(Totalitems, len(items))
	}
	return items, paginationResult, nil
}
