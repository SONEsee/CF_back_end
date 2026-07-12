package requestbody

type ProductVariantRequestBody struct {
	ProductID   int     `json:"product_id" validate:"required,gt=0"`
	VariantName string  `json:"variant_name" validate:"required,min=1,max=100"`
	SkuCode     string  `json:"sku_code" validate:"required,max=50"`
	CfCode      string  `json:"cf_code" validate:"omitempty,max=30"`
	Barcode     string  `json:"barcode" validate:"omitempty,max=50"`
	Price       float64 `json:"price" validate:"required,gte=0"`
	CostPrice   float64 `json:"cost_price" validate:"omitempty,gte=0"`
	WeightGrams int     `json:"weight_grams" validate:"omitempty,gte=0"`
}

type ProductVariantPatchRequest struct {
	VariantName *string  `json:"variant_name,omitempty" validate:"omitempty,min=1,max=100"`
	SkuCode     *string  `json:"sku_code,omitempty" validate:"omitempty,max=50"`
	CfCode      *string  `json:"cf_code,omitempty" validate:"omitempty,max=30"`
	Barcode     *string  `json:"barcode,omitempty" validate:"omitempty,max=50"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gte=0"`
	CostPrice   *float64 `json:"cost_price,omitempty" validate:"omitempty,gte=0"`
	WeightGrams *int     `json:"weight_grams,omitempty" validate:"omitempty,gte=0"`
	IsActive    *bool    `json:"is_active,omitempty"`
}
