package requestbody

type UserRequestBody struct {
	ShopID       *int   `json:"shop_id" validate:"omitempty,gt=0"`
	RoleID       int    `json:"role_id" validate:"required,gt=0"`
	Username     string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password     string `json:"password" validate:"required,min=8"`
	FullName     string `json:"full_name" validate:"required,min=2,max=100"`
	Email        string `json:"email" validate:"omitempty,email"`
	Phone        string `json:"phone" validate:"omitempty,max=20"`
	ProfileImage string `json:"profile_image" validate:"omitempty,max=500"`
}

type UserRequestBodyPacth struct {
	ShopID       *int    `json:"shop_id,omitempty" validate:"omitempty,gt=0"`
	RoleID       *int    `json:"role_id,omitempty" validate:"omitempty,gt=0"`
	Username     *string `json:"username,omitempty" validate:"omitempty,min=3,max=50,alphanum"`
	Password     string  `json:"password,omitempty" validate:"omitempty,min=8"`
	FullName     *string `json:"full_name,omitempty" validate:"omitempty,min=2,max=100"`
	Email        *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone        *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	ProfileImage *string `json:"profile_image,omitempty" validate:"omitempty,max=500"`
}
