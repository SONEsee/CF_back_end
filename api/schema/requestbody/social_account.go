package requestbody

type SocialAccountRequestBody struct {
	ShopID            int     `json:"shop_id" validate:"required,gt=0"`
	Platform          string  `json:"platform" validate:"required,oneof=FACEBOOK_PAGE INSTAGRAM LINE_OA"`
	PlatformAccountID string  `json:"platform_account_id" validate:"required,max=100"`
	AccountName       string  `json:"account_name" validate:"omitempty,max=150"`
	AccessToken       string  `json:"access_token" validate:"omitempty"`
	TokenExpiresAt    *string `json:"token_expires_at" validate:"omitempty"`
}

type SocialAccountPatchRequest struct {
	AccountName    *string `json:"account_name,omitempty" validate:"omitempty,max=150"`
	AccessToken    *string `json:"access_token,omitempty"`
	TokenExpiresAt *string `json:"token_expires_at,omitempty"`
}
