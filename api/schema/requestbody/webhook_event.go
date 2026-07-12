package requestbody

type WebhookEventRequestBody struct {
	SocialAccountID int    `json:"social_account_id" validate:"required,gt=0"`
	EventType       string `json:"event_type" validate:"omitempty,max=50"`
	RawPayload      string `json:"raw_payload" validate:"omitempty,json"`
}
