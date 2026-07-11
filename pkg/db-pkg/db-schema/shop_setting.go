package dbschema

import "time"

type ShopSettingDBSchema struct {
	ID            int       `db:"id" json:"id"`
	ShopID        int       `db:"shop_id" json:"shop_id"`
	Currency      string    `db:"currency" json:"currency"`
	VatRate       float64   `db:"vat_rate" json:"vat_rate"`
	AutoReplyMsg  string    `db:"auto_reply_msg" json:"auto_reply_msg"`
	BusinessHours string    `db:"business_hours" json:"business_hours"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
