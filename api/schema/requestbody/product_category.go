package requestbody

type ProductCategoryRequestBody struct {
	ShopID    int    `json:"shop_id" validate:"required,gt=0"`
	ParentID  *int   `json:"parent_id" validate:"omitempty,gt=0"`
	Name      string `json:"name" validate:"required,min=2,max=100"`
	SortOrder int    `json:"sort_order"`
}

type ProductCategoryPatchRequest struct {
	ParentID  *int    `json:"parent_id,omitempty" validate:"omitempty,gt=0"`
	Name      *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	SortOrder *int    `json:"sort_order,omitempty"`
}
