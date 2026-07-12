package requestbody

type LiveSessionProductRequestBody struct {
	LiveSessionID    int      `json:"live_session_id" validate:"required,gt=0"`
	ProductVariantID int      `json:"product_variant_id" validate:"required,gt=0"`
	LivePrice        *float64 `json:"live_price" validate:"omitempty,gte=0"`
	CfCodeOverride   string   `json:"cf_code_override" validate:"omitempty,max=30"`
}

type LiveSessionProductPatchRequest struct {
	LivePrice      *float64 `json:"live_price,omitempty" validate:"omitempty,gte=0"`
	CfCodeOverride *string  `json:"cf_code_override,omitempty" validate:"omitempty,max=30"`
}
