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

var shipmentColumns = []string{"id", "order_id", "courier_name", "tracking_number", "label_pdf_url", "shipping_status", "shipped_at", "delivered_at"}

func scanShipment(row pgx.Row, item *dbschema.ShipmentDBSchema) error {
	return row.Scan(
		&item.ID, &item.OrderID, &item.CourierName, &item.TrackingNumber,
		&item.LabelPdfURL, &item.ShippingStatus, &item.ShippedAt, &item.DeliveredAt,
	)
}

func GetShipmentDataQuery(ctx context.Context, id *int, orderID *int, paginationParams *pagination.PaginationParams) ([]dbschema.ShipmentDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(shipmentColumns...).From(`"shipments"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ShipmentDBSchema
		if err := scanShipment(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("shipment with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ShipmentDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"shipments"`)
	listQuery := psql.Select(shipmentColumns...).From(`"shipments"`)
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
		return nil, nil, fmt.Errorf("failed to query shipments: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ShipmentDBSchema
	for rows.Next() {
		var item dbschema.ShipmentDBSchema
		if err := scanShipment(rows, &item); err != nil {
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
