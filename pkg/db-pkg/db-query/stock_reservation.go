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

var stockReservationColumns = []string{"id", "product_variant_id", "customer_id", "order_item_id", "reserved_qty", "expires_at", "status", "created_at"}

func scanStockReservation(row pgx.Row, item *dbschema.StockReservationDBSchema) error {
	return row.Scan(
		&item.ID, &item.ProductVariantID, &item.CustomerID, &item.OrderItemID,
		&item.ReservedQty, &item.ExpiresAt, &item.Status, &item.CreatedAt,
	)
}

func GetStockReservationDataQuery(ctx context.Context, id *int, productVariantID *int, paginationParams *pagination.PaginationParams) ([]dbschema.StockReservationDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(stockReservationColumns...).From(`"stock_reservations"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.StockReservationDBSchema
		if err := scanStockReservation(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("stock reservation with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.StockReservationDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"stock_reservations"`)
	listQuery := psql.Select(stockReservationColumns...).From(`"stock_reservations"`)
	if productVariantID != nil {
		countQuery = countQuery.Where("product_variant_id=?", *productVariantID)
		listQuery = listQuery.Where("product_variant_id=?", *productVariantID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("id DESC")
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
		return nil, nil, fmt.Errorf("failed to query stock reservations: %w", err)
	}
	defer rows.Close()

	var items []dbschema.StockReservationDBSchema
	for rows.Next() {
		var item dbschema.StockReservationDBSchema
		if err := scanStockReservation(rows, &item); err != nil {
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

// GetStockReservationByOrderItemIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນປ່ຽນ order status (PAID/CANCELLED) ໃນ transaction ດຽວກັນ
func GetStockReservationByOrderItemIDForUpdate(ctx context.Context, tx dbpkg.DBTX, orderItemID int) (*dbschema.StockReservationDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(stockReservationColumns...).From(`"stock_reservations"`).Where("order_item_id=?", orderItemID).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.StockReservationDBSchema
	if err := scanStockReservation(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("stock reservation for order_item_id %d not found", orderItemID)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}

// GetStockReservationByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນປ່ຽນ status (COMPLETED/EXPIRED) ໃນ transaction ດຽວກັບການອັບເດດ inventories
func GetStockReservationByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int64) (*dbschema.StockReservationDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(stockReservationColumns...).From(`"stock_reservations"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.StockReservationDBSchema
	if err := scanStockReservation(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("stock reservation with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
