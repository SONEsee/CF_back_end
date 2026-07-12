package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateOrder ສ້າງ order (current_status ໃຊ້ default UNPAID ຂອງ DB) ແລະ ຄືນ id ໃໝ່
func CreateOrder(ctx context.Context, tx dbpkg.DBTX, shopID, customerID int, liveSessionID, discountID *int, orderNumber string, itemsTotal, discountAmount, shippingFee, netPayable float64, note string) (int64, error) {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"orders"`).
		Columns("shop_id", "customer_id", "live_session_id", "discount_id", "order_number", "items_total_amount", "discount_amount", "shipping_fee", "net_payable_amount", "note").
		Values(shopID, customerID, liveSessionID, discountID, orderNumber, itemsTotal, discountAmount, shippingFee, netPayable, nullableStr(note)).
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
