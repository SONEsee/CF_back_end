package dbschema

import "time"

type OrderDBSchema struct {
	ID               int       `db:"id" json:"id"`
	ShopID           int       `db:"shop_id" json:"shop_id"`
	CustomerID       int       `db:"customer_id" json:"customer_id"`
	LiveSessionID    *int      `db:"live_session_id" json:"live_session_id"`
	DiscountID       *int      `db:"discount_id" json:"discount_id"`
	OrderNumber      string    `db:"order_number" json:"order_number"`
	CurrentStatus    string    `db:"current_status" json:"current_status"`
	ItemsTotalAmount float64   `db:"items_total_amount" json:"items_total_amount"`
	DiscountAmount   float64   `db:"discount_amount" json:"discount_amount"`
	ShippingFee      float64   `db:"shipping_fee" json:"shipping_fee"`
	NetPayableAmount float64   `db:"net_payable_amount" json:"net_payable_amount"`
	Note             *string   `db:"note" json:"note"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
