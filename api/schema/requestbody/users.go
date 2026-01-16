package requestbody

type UserRequestBody struct {
	Name       string `json:"name" validate:"required,min=2,max=100"`
	FullName   string `json:"full_name" validate:"required,min=2,max=200"`
	UserName   string `json:"user_name" validate:"required,min=3,max=50,alphanum"`
	Password   string `json:"password" validate:"required,min=8"`
	ProfileImg string `json:"profile_image" validate:"omitempty,url"`
	BlackList  bool   `json:"black_list"`
	RoleID     int    `json:"role_id" validate:"required,gt=0"`
}
type UserRequestBodyPacth struct {
	Name       *string `json:"name,omitempty" validate:"omitempty,min=1,max=150"`
	FullName   *string `json:"full_name,omitempty" validate:"omitempty,min=1,max=150"`
	UserName   *string `json:"user_name,omitempty" validate:"omitempty,min=1,max=150"`
	Password   string  `json:"password" validate:"omitempty,min=8"`
	ProfileImg *string `json:"profile_image" validate:"omitempty,url"`
	BlackList  *string `json:"black_list,omitempty"`
	RoleID     *int    `json:"role_id" validate:"omitempty,gt=0"`
}
