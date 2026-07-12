package dbschema

import "time"

type RefundDBSchema struct {
	ID           int       `db:"id" json:"id"`
	OrderID      int       `db:"order_id" json:"order_id"`
	Reason       *string   `db:"reason" json:"reason"`
	RefundAmount float64   `db:"refund_amount" json:"refund_amount"`
	Status       string    `db:"status" json:"status"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
