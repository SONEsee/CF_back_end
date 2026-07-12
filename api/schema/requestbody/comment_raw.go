package requestbody

type CommentRawRequestBody struct {
	LiveSessionID  int    `json:"live_session_id" validate:"required,gt=0"`
	FbCommentID    string `json:"fb_comment_id" validate:"omitempty,max=100"`
	FbUserID       string `json:"fb_user_id" validate:"omitempty,max=100"`
	CommentMessage string `json:"comment_message"`
}
