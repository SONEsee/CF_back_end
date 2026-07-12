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

var discountColumns = []string{"id", "shop_id", "code", "discount_type", "discount_value", "min_order", "usage_limit", "used_count", "start_at", "end_at", "is_active"}

func scanDiscount(row pgx.Row, item *dbschema.DiscountDBSchema) error {
	return row.Scan(
		&item.ID, &item.ShopID, &item.Code, &item.DiscountType, &item.DiscountValue,
		&item.MinOrder, &item.UsageLimit, &item.UsedCount, &item.StartAt, &item.EndAt, &item.IsActive,
	)
}

func GetDiscountDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.DiscountDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(discountColumns...).From(`"discounts"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.DiscountDBSchema
		if err := scanDiscount(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("discount with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.DiscountDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"discounts"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(discountColumns...).From(`"discounts"`).OrderBy("id ASC")

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
		return nil, nil, fmt.Errorf("failed to query discounts: %w", err)
	}
	defer rows.Close()

	var items []dbschema.DiscountDBSchema
	for rows.Next() {
		var item dbschema.DiscountDBSchema
		if err := scanDiscount(rows, &item); err != nil {
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

// GetDiscountByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນສ້າງ order ເພື່ອ validate+increment used_count ໃນ transaction ດຽວ
func GetDiscountByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int) (*dbschema.DiscountDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(discountColumns...).From(`"discounts"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.DiscountDBSchema
	if err := scanDiscount(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("discount with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
