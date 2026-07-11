package requestbody

type ShopBankAccountRequestBody struct {
	ShopID        int    `json:"shop_id" validate:"required,gt=0"`
	BankName      string `json:"bank_name" validate:"required,min=2,max=50"`
	AccountNumber string `json:"account_number" validate:"required,max=30"`
	AccountName   string `json:"account_name" validate:"required,max=150"`
	PromptpayID   string `json:"promptpay_id" validate:"omitempty,max=30"`
}

type ShopBankAccountPatchRequest struct {
	BankName      *string `json:"bank_name,omitempty" validate:"omitempty,min=2,max=50"`
	AccountNumber *string `json:"account_number,omitempty" validate:"omitempty,max=30"`
	AccountName   *string `json:"account_name,omitempty" validate:"omitempty,max=150"`
	PromptpayID   *string `json:"promptpay_id,omitempty" validate:"omitempty,max=30"`
	IsActive      *bool   `json:"is_active,omitempty"`
}
