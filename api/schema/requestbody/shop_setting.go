package requestbody

type ShopSettingRequestBody struct {
	ShopID        int     `json:"shop_id" validate:"required,gt=0"`
	Currency      string  `json:"currency" validate:"omitempty,len=3"`
	VatRate       float64 `json:"vat_rate" validate:"omitempty,gte=0"`
	AutoReplyMsg  string  `json:"auto_reply_msg"`
	BusinessHours string  `json:"business_hours" validate:"omitempty,json"`
}

type ShopSettingPatchRequest struct {
	Currency      *string  `json:"currency,omitempty" validate:"omitempty,len=3"`
	VatRate       *float64 `json:"vat_rate,omitempty" validate:"omitempty,gte=0"`
	AutoReplyMsg  *string  `json:"auto_reply_msg,omitempty"`
	BusinessHours *string  `json:"business_hours,omitempty" validate:"omitempty,json"`
}
