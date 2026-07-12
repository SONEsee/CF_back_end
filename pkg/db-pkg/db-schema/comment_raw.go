package dbschema

import "time"

type CommentRawDBSchema struct {
	ID             int64     `db:"id" json:"id"`
	LiveSessionID  int       `db:"live_session_id" json:"live_session_id"`
	FbCommentID    *string   `db:"fb_comment_id" json:"fb_comment_id"`
	FbUserID       *string   `db:"fb_user_id" json:"fb_user_id"`
	CommentMessage *string   `db:"comment_message" json:"comment_message"`
	IsProcessed    bool      `db:"is_processed" json:"is_processed"`
	ReceivedAt     time.Time `db:"received_at" json:"received_at"`
}
