package dbschema

import "time"

type DiscountDBSchema struct {
	ID            int        `db:"id" json:"id"`
	ShopID        int        `db:"shop_id" json:"shop_id"`
	Code          string     `db:"code" json:"code"`
	DiscountType  string     `db:"discount_type" json:"discount_type"`
	DiscountValue float64    `db:"discount_value" json:"discount_value"`
	MinOrder      float64    `db:"min_order" json:"min_order"`
	UsageLimit    *int       `db:"usage_limit" json:"usage_limit"`
	UsedCount     int        `db:"used_count" json:"used_count"`
	StartAt       *time.Time `db:"start_at" json:"start_at"`
	EndAt         *time.Time `db:"end_at" json:"end_at"`
	IsActive      bool       `db:"is_active" json:"is_active"`
}
