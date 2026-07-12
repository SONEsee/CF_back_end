package requestbody

type LiveSessionRequestBody struct {
	SocialAccountID int    `json:"social_account_id" validate:"required,gt=0"`
	FbVideoID       string `json:"fb_video_id" validate:"omitempty,max=100"`
	SessionTitle    string `json:"session_title" validate:"omitempty,max=150"`
}

type LiveSessionPatchRequest struct {
	SessionTitle *string `json:"session_title,omitempty" validate:"omitempty,max=150"`
}
