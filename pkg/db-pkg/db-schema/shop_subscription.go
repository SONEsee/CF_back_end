package dbschema

import "time"

type ShopSubscriptionDBSchema struct {
	ID        int        `db:"id" json:"id"`
	ShopID    int        `db:"shop_id" json:"shop_id"`
	PlanID    int        `db:"plan_id" json:"plan_id"`
	StartDate time.Time  `db:"start_date" json:"start_date"`
	EndDate   *time.Time `db:"end_date" json:"end_date"`
	Status    string     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}
