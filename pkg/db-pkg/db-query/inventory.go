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

var inventoryColumns = []string{"id", "product_variant_id", "actual_qty", "reserved_qty", "available_qty", "reorder_level", "last_updated"}

func scanInventory(row pgx.Row, item *dbschema.InventoryDBSchema) error {
	return row.Scan(
		&item.ID, &item.ProductVariantID, &item.ActualQty, &item.ReservedQty,
		&item.AvailableQty, &item.ReorderLevel, &item.LastUpdated,
	)
}

func GetInventoryDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.InventoryDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(inventoryColumns...).From(`"inventories"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.InventoryDBSchema
		if err := scanInventory(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("inventory with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.InventoryDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"inventories"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(inventoryColumns...).From(`"inventories"`).OrderBy("id ASC")

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
		return nil, nil, fmt.Errorf("failed to query inventories: %w", err)
	}
	defer rows.Close()

	var items []dbschema.InventoryDBSchema
	for rows.Next() {
		var item dbschema.InventoryDBSchema
		if err := scanInventory(rows, &item); err != nil {
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

// GetInventoryByVariantIDForUpdate ອ່ານ inventory ພ້ອມລັອກແຖວ (SELECT ... FOR UPDATE) — ໃຊ້ພາຍໃນ transaction ດຽວກັບການສ້າງ stock_movement ເພື່ອປ້ອງກັນ race condition ຕອນອັບເດດ qty ພ້ອມກັນ
func GetInventoryByVariantIDForUpdate(ctx context.Context, tx dbpkg.DBTX, productVariantID int) (*dbschema.InventoryDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(inventoryColumns...).From(`"inventories"`).Where("product_variant_id=?", productVariantID).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.InventoryDBSchema
	if err := scanInventory(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("inventory for product_variant_id %d not found", productVariantID)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
