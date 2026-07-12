package requestbody

type ParseCommentRequest struct {
	CommentRawID int64 `json:"comment_raw_id" validate:"required,gt=0"`
}
