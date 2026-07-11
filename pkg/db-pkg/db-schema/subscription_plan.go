package dbschema

type SubscriptionPlanDBSchema struct {
	ID           int     `db:"id" json:"id"`
	PlanName     string  `db:"plan_name" json:"plan_name"`
	PriceMonthly float64 `db:"price_monthly" json:"price_monthly"`
	MaxUsers     int     `db:"max_users" json:"max_users"`
	MaxProducts  int     `db:"max_products" json:"max_products"`
	Features     string  `db:"features" json:"features"`
}
