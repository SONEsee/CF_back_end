package requestbody

type SubscriptionPlanRequestBody struct {
	PlanName     string  `json:"plan_name" validate:"required,min=2,max=50"`
	PriceMonthly float64 `json:"price_monthly" validate:"required,gte=0"`
	MaxUsers     int     `json:"max_users" validate:"required,gt=0"`
	MaxProducts  int     `json:"max_products" validate:"required,gt=0"`
	Features     string  `json:"features" validate:"omitempty,json"`
}

type SubscriptionPlanPatchRequest struct {
	PlanName     *string  `json:"plan_name,omitempty" validate:"omitempty,min=2,max=50"`
	PriceMonthly *float64 `json:"price_monthly,omitempty" validate:"omitempty,gte=0"`
	MaxUsers     *int     `json:"max_users,omitempty" validate:"omitempty,gt=0"`
	MaxProducts  *int     `json:"max_products,omitempty" validate:"omitempty,gt=0"`
	Features     *string  `json:"features,omitempty" validate:"omitempty,json"`
}
