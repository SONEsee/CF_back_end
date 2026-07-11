package dbschema

type ShopBankAccountDBSchema struct {
	ID            int    `db:"id" json:"id"`
	ShopID        int    `db:"shop_id" json:"shop_id"`
	BankName      string `db:"bank_name" json:"bank_name"`
	AccountNumber string `db:"account_number" json:"account_number"`
	AccountName   string `db:"account_name" json:"account_name"`
	PromptpayID   string `db:"promptpay_id" json:"promptpay_id"`
	IsActive      bool   `db:"is_active" json:"is_active"`
}
