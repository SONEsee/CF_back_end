package requestbody

type UserLoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	UserName   string `json:"user_name"`
	ProfileImg string `json:"profile_image,omitempty"`
	RoleID     int    `json:"role_id"`
	Token      string `json:"token,omitempty"`
}
