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

var orderColumns = []string{"id", "shop_id", "customer_id", "live_session_id", "discount_id", "order_number", "current_status", "items_total_amount", "discount_amount", "shipping_fee", "net_payable_amount", "note", "created_at", "updated_at"}

func scanOrder(row pgx.Row, item *dbschema.OrderDBSchema) error {
	return row.Scan(
		&item.ID, &item.ShopID, &item.CustomerID, &item.LiveSessionID, &item.DiscountID, &item.OrderNumber,
		&item.CurrentStatus, &item.ItemsTotalAmount, &item.DiscountAmount, &item.ShippingFee,
		&item.NetPayableAmount, &item.Note, &item.CreatedAt, &item.UpdatedAt,
	)
}

func GetOrderDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.OrderDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(orderColumns...).From(`"orders"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.OrderDBSchema
		if err := scanOrder(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("order with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.OrderDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"orders"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(orderColumns...).From(`"orders"`).OrderBy("id DESC")

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
		return nil, nil, fmt.Errorf("failed to query orders: %w", err)
	}
	defer rows.Close()

	var items []dbschema.OrderDBSchema
	for rows.Next() {
		var item dbschema.OrderDBSchema
		if err := scanOrder(rows, &item); err != nil {
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

// GetOrderByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນປ່ຽນ status (Step 3)
func GetOrderByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int64) (*dbschema.OrderDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(orderColumns...).From(`"orders"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.OrderDBSchema
	if err := scanOrder(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("order with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
