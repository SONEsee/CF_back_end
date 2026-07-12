package dbschema

import "time"

type ChatMessageDBSchema struct {
	ID             int64     `db:"id" json:"id"`
	ConversationID int       `db:"conversation_id" json:"conversation_id"`
	SenderType     string    `db:"sender_type" json:"sender_type"`
	MessageType    string    `db:"message_type" json:"message_type"`
	MessageBody    *string   `db:"message_body" json:"message_body"`
	AttachmentURL  *string   `db:"attachment_url" json:"attachment_url"`
	IsRead         bool      `db:"is_read" json:"is_read"`
	SentAt         time.Time `db:"sent_at" json:"sent_at"`
}
