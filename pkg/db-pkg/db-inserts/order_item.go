package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateOrderItem ສ້າງລາຍການສິນຄ້າ (priceSnapshot/subtotal ຄິດໄລ່ມາຈາກ service ແລ້ວ) ແລະ ຄືນ id ໃໝ່
func CreateOrderItem(ctx context.Context, tx dbpkg.DBTX, orderID, productVariantID, buyQuantity int, priceSnapshot, subtotal float64) (int64, error) {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"order_items"`).
		Columns("order_id", "product_variant_id", "buy_quantity", "price_snapshot", "subtotal").
		Values(orderID, productVariantID, buyQuantity, priceSnapshot, subtotal).
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
