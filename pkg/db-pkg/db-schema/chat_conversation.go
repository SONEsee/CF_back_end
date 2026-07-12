package dbschema

import "time"

type ChatConversationDBSchema struct {
	ID                 int       `db:"id" json:"id"`
	SocialAccountID    int       `db:"social_account_id" json:"social_account_id"`
	CustomerID         int       `db:"customer_id" json:"customer_id"`
	AssignedStaffID    *int      `db:"assigned_staff_id" json:"assigned_staff_id"`
	LastMessagePreview *string   `db:"last_message_preview" json:"last_message_preview"`
	UnreadCount        int       `db:"unread_count" json:"unread_count"`
	Status             string    `db:"status" json:"status"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
}
