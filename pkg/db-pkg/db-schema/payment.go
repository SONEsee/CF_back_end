package dbschema

import "time"

type PaymentDBSchema struct {
	ID                int        `db:"id" json:"id"`
	OrderID           int        `db:"order_id" json:"order_id"`
	ShopBankAccountID *int       `db:"shop_bank_account_id" json:"shop_bank_account_id"`
	PaymentMethod     string     `db:"payment_method" json:"payment_method"`
	SlipImagePath     *string    `db:"slip_image_path" json:"slip_image_path"`
	BankTransRefID    *string    `db:"bank_trans_ref_id" json:"bank_trans_ref_id"`
	VerifiedAmount    *float64   `db:"verified_amount" json:"verified_amount"`
	IsValidSlip       *bool      `db:"is_valid_slip" json:"is_valid_slip"`
	PaidAt            *time.Time `db:"paid_at" json:"paid_at"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
}
