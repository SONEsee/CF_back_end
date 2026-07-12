package requestbody

type DiscountRequestBody struct {
	ShopID        int     `json:"shop_id" validate:"required,gt=0"`
	Code          string  `json:"code" validate:"required,min=2,max=50"`
	DiscountType  string  `json:"discount_type" validate:"required,oneof=PERCENT FIXED"`
	DiscountValue float64 `json:"discount_value" validate:"required,gt=0"`
	MinOrder      float64 `json:"min_order" validate:"omitempty,gte=0"`
	UsageLimit    *int    `json:"usage_limit" validate:"omitempty,gt=0"`
	StartAt       *string `json:"start_at" validate:"omitempty"`
	EndAt         *string `json:"end_at" validate:"omitempty"`
}

type DiscountPatchRequest struct {
	DiscountValue *float64 `json:"discount_value,omitempty" validate:"omitempty,gt=0"`
	MinOrder      *float64 `json:"min_order,omitempty" validate:"omitempty,gte=0"`
	UsageLimit    *int     `json:"usage_limit,omitempty" validate:"omitempty,gt=0"`
	StartAt       *string  `json:"start_at,omitempty"`
	EndAt         *string  `json:"end_at,omitempty"`
}
