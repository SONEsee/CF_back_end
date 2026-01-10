package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	"github.com/jackc/pgx/v5"

	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

func GetDataTypeMidsine(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.TypeMedicineDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()
	if id != nil {
		// ✅ ເພີ່ມ id_type ໃນ SELECT
		query := psql.Select("id_type", "name_type", "detail_type").From(`"TypeMidisine"`).Where("id_type=?", *id).Where("deleted_at IS NULL")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed for convert sql %w", err)
		}
		var item dbschema.TypeMedicineDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID, // ✅ ເພີ່ມ scan id_type
			&item.NameType,
			&item.Detail_Type,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("type medicine with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("fail get type medicine %w", err)
		}
		return []dbschema.TypeMedicineDBSchema{item}, nil, nil
	}

	var totalItem int
	querCount := psql.Select("COUNT(*)").From(`"TypeMidisine"`).Where("deleted_at IS NULL")
	sqlCount, argsCount, err := querCount.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed for convert sql %w", err)
	}
	err = dbpkg.DB.QueryRow(ctx, sqlCount, argsCount...).Scan(&totalItem)
	if err != nil {
		return nil, nil, fmt.Errorf("failed count %w", err)
	}

	// ✅ ເພີ່ມ id_type ໃນ SELECT ແລະໃຊ້ id_type ໃນ ORDER BY
	query := psql.Select("id_type", "name_type", "detail_type").From(`"TypeMidisine"`).Where("deleted_at IS NULL").OrderBy("id_type ASC")
	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		query = query.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed convert for sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed execute %w", err)
	}
	defer rows.Close()

	var items []dbschema.TypeMedicineDBSchema
	for rows.Next() {
		var item dbschema.TypeMedicineDBSchema
		err := rows.Scan(
			&item.ID, // ✅ ເພີ່ມ scan id_type
			&item.NameType,
			&item.Detail_Type,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed scan data %w", err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("iterating rows: %w", err)
	}
	if paginationParams != nil && paginationParams.IsValid() {
		paginationResult = paginationParams.CalculatePagination(totalItem, len(items))
	}
	return items, paginationResult, nil
}
