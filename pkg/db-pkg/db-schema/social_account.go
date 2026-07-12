package dbschema

import "time"

type SocialAccountDBSchema struct {
	ID                int        `db:"id" json:"id"`
	ShopID            int        `db:"shop_id" json:"shop_id"`
	Platform          string     `db:"platform" json:"platform"`
	PlatformAccountID string     `db:"platform_account_id" json:"platform_account_id"`
	AccountName       *string    `db:"account_name" json:"account_name"`
	AccessToken       *string    `db:"access_token" json:"-"`
	TokenExpiresAt    *time.Time `db:"token_expires_at" json:"token_expires_at"`
	IsActive          bool       `db:"is_active" json:"is_active"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
}
