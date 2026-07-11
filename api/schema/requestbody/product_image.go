package requestbody

type ProductImageRequestBody struct {
	ProductID int    `json:"product_id" validate:"required,gt=0"`
	ImageURL  string `json:"image_url" validate:"required,url"`
	SortOrder int    `json:"sort_order"`
}

type ProductImagePatchRequest struct {
	ImageURL  *string `json:"image_url,omitempty" validate:"omitempty,url"`
	SortOrder *int    `json:"sort_order,omitempty"`
}
