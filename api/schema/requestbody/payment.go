package requestbody

type PaymentRequestBody struct {
	OrderID           int    `json:"order_id" validate:"required,gt=0"`
	ShopBankAccountID *int   `json:"shop_bank_account_id" validate:"omitempty,gt=0"`
	PaymentMethod     string `json:"payment_method" validate:"required,oneof=SLIP PROMPTPAY COD"`
	SlipImagePath     string `json:"slip_image_path" validate:"omitempty,max=255"`
	BankTransRefID    string `json:"bank_trans_ref_id" validate:"omitempty,max=100"`
}

// PaymentVerifyRequest — staff ຢືນຢັນສະລິບ: ຖືກຕ້ອງ+ຍອດຈິງ, ຫຼືປະຕິເສດ
type PaymentVerifyRequest struct {
	IsValidSlip    bool    `json:"is_valid_slip"`
	VerifiedAmount float64 `json:"verified_amount" validate:"omitempty,gte=0"`
}
