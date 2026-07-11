package requestbody

type RoleRequestBody struct {
	ShopID      *int   `json:"shop_id" validate:"omitempty,gt=0"`
	RoleName    string `json:"role_name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type RolePatchRequest struct {
	ShopID      *int    `json:"shop_id,omitempty" validate:"omitempty,gt=0"`
	RoleName    *string `json:"role_name,omitempty" validate:"omitempty,min=2,max=50"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=255"`
}
