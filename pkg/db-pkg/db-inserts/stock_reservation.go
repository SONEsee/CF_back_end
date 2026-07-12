package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateStockReservation ສ້າງແຖວ (status ໃຊ້ default HOLDING ຂອງ DB) ແລະ ຄືນ id ໃໝ່
func CreateStockReservation(ctx context.Context, tx dbpkg.DBTX, req requestbody.StockReservationRequestBody) (int64, error) {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"stock_reservations"`).
		Columns("product_variant_id", "customer_id", "order_item_id", "reserved_qty", "expires_at").
		Values(req.ProductVariantID, req.CustomerID, req.OrderItemID, req.ReservedQty, req.ExpiresAt).
		Suffix("RETURNING id")
	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	if err := tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
