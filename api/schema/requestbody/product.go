package requestbody

type ProductRequestBody struct {
	ShopID       int    `json:"shop_id" validate:"required,gt=0"`
	CategoryID   *int   `json:"category_id" validate:"omitempty,gt=0"`
	ProductName  string `json:"product_name" validate:"required,min=2,max=150"`
	Description  string `json:"description"`
	ImageMainURL string `json:"image_main_url" validate:"omitempty,url"`
}

type ProductPatchRequest struct {
	CategoryID   *int    `json:"category_id,omitempty" validate:"omitempty,gt=0"`
	ProductName  *string `json:"product_name,omitempty" validate:"omitempty,min=2,max=150"`
	Description  *string `json:"description,omitempty"`
	ImageMainURL *string `json:"image_main_url,omitempty" validate:"omitempty,url"`
	IsActive     *bool   `json:"is_active,omitempty"`
}
