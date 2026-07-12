package requestbody

// SendChatMessageRequest — ບໍ່ຮັບ conversation_id ໂດຍກົງ, server ຈະ find-or-create conversation ໃຫ້ຈາກ (social_account_id, customer_id)
type SendChatMessageRequest struct {
	SocialAccountID int    `json:"social_account_id" validate:"required,gt=0"`
	CustomerID      int    `json:"customer_id" validate:"required,gt=0"`
	SenderType      string `json:"sender_type" validate:"required,oneof=CUSTOMER BOT_SYSTEM STAFF_AGENT"`
	MessageType     string `json:"message_type" validate:"omitempty,oneof=TEXT IMAGE PAYMENT_LINK"`
	MessageBody     string `json:"message_body"`
	AttachmentURL   string `json:"attachment_url" validate:"omitempty,url"`
}
