package dbschema

import "time"

type OrderStatusLogDBSchema struct {
	ID            int       `db:"id" json:"id"`
	OrderID       int       `db:"order_id" json:"order_id"`
	FromStatus    *string   `db:"from_status" json:"from_status"`
	ToStatus      string    `db:"to_status" json:"to_status"`
	ChangedByType string    `db:"changed_by_type" json:"changed_by_type"`
	ChangedByID   *int      `db:"changed_by_id" json:"changed_by_id"`
	Note          *string   `db:"note" json:"note"`
	ChangedAt     time.Time `db:"changed_at" json:"changed_at"`
}
