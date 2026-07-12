package dbschema

import "time"

type StockReservationDBSchema struct {
	ID               int       `db:"id" json:"id"`
	ProductVariantID int       `db:"product_variant_id" json:"product_variant_id"`
	CustomerID       *int      `db:"customer_id" json:"customer_id"`
	OrderItemID      *int      `db:"order_item_id" json:"order_item_id"`
	ReservedQty      int       `db:"reserved_qty" json:"reserved_qty"`
	ExpiresAt        time.Time `db:"expires_at" json:"expires_at"`
	Status           string    `db:"status" json:"status"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
