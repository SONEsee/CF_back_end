package dbschema

type OrderItemDBSchema struct {
	ID               int     `db:"id" json:"id"`
	OrderID          int     `db:"order_id" json:"order_id"`
	ProductVariantID int     `db:"product_variant_id" json:"product_variant_id"`
	BuyQuantity      int     `db:"buy_quantity" json:"buy_quantity"`
	PriceSnapshot    float64 `db:"price_snapshot" json:"price_snapshot"`
	Subtotal         float64 `db:"subtotal" json:"subtotal"`
}
