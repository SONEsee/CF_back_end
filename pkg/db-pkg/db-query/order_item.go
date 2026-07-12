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

var orderItemColumns = []string{"id", "order_id", "product_variant_id", "buy_quantity", "price_snapshot", "subtotal"}

func scanOrderItem(row pgx.Row, item *dbschema.OrderItemDBSchema) error {
	return row.Scan(&item.ID, &item.OrderID, &item.ProductVariantID, &item.BuyQuantity, &item.PriceSnapshot, &item.Subtotal)
}

// GetOrderItemDataQuery — id ຫາລາຍການດຽວ, ຫຼືກັ່ນຕອງດ້ວຍ orderID ເບິ່ງລາຍການທັງໝົດຂອງ order ນັ້ນ
func GetOrderItemDataQuery(ctx context.Context, id *int, orderID *int, paginationParams *pagination.PaginationParams) ([]dbschema.OrderItemDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(orderItemColumns...).From(`"order_items"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.OrderItemDBSchema
		if err := scanOrderItem(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("order item with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.OrderItemDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"order_items"`)
	listQuery := psql.Select(orderItemColumns...).From(`"order_items"`)
	if orderID != nil {
		countQuery = countQuery.Where("order_id=?", *orderID)
		listQuery = listQuery.Where("order_id=?", *orderID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("id ASC")
	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		listQuery = listQuery.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := listQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query order items: %w", err)
	}
	defer rows.Close()

	var items []dbschema.OrderItemDBSchema
	for rows.Next() {
		var item dbschema.OrderItemDBSchema
		if err := scanOrderItem(rows, &item); err != nil {
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
