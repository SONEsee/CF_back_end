package dbschema

import "time"

type WebhookEventDBSchema struct {
	ID              int       `db:"id" json:"id"`
	SocialAccountID int       `db:"social_account_id" json:"social_account_id"`
	EventType       *string   `db:"event_type" json:"event_type"`
	RawPayload      *string   `db:"raw_payload" json:"raw_payload"`
	Processed       bool      `db:"processed" json:"processed"`
	ReceivedAt      time.Time `db:"received_at" json:"received_at"` 
}
