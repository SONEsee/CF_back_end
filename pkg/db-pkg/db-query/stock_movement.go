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

var stockMovementColumns = []string{"id", "product_variant_id", "movement_type", "qty_change", "balance_after", "ref_type", "ref_id", "note", "created_by", "created_at"}

func scanStockMovement(row pgx.Row, item *dbschema.StockMovementDBSchema) error {
	return row.Scan(
		&item.ID, &item.ProductVariantID, &item.MovementType, &item.QtyChange, &item.BalanceAfter,
		&item.RefType, &item.RefID, &item.Note, &item.CreatedBy, &item.CreatedAt,
	)
}

// GetStockMovementDataQuery — id ຫາລາຍການດຽວ, ຫຼືກັ່ນຕອງດ້ວຍ productVariantID ເບິ່ງ ledger ຂອງ variant ນັ້ນ (ຖ້າບໍ່ໃສ່ທັງສອງ ຈະໄດ້ທັງໝົດ)
func GetStockMovementDataQuery(ctx context.Context, id *int64, productVariantID *int, paginationParams *pagination.PaginationParams) ([]dbschema.StockMovementDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(stockMovementColumns...).From(`"stock_movements"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.StockMovementDBSchema
		if err := scanStockMovement(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("stock movement with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.StockMovementDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"stock_movements"`)
	listQuery := psql.Select(stockMovementColumns...).From(`"stock_movements"`)
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
		return nil, nil, fmt.Errorf("failed to query stock movements: %w", err)
	}
	defer rows.Close()

	var items []dbschema.StockMovementDBSchema
	for rows.Next() {
		var item dbschema.StockMovementDBSchema
		if err := scanStockMovement(rows, &item); err != nil {
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
