package dbschema

import "time"

type CustomerDBSchema struct {
	ID               int       `db:"id" json:"id"`
	ShopID           int       `db:"shop_id" json:"shop_id"`
	SocialPlatformID *string   `db:"social_platform_id" json:"social_platform_id"`
	CustomerName     *string   `db:"customer_name" json:"customer_name"`
	ProfilePicURL    *string   `db:"profile_pic_url" json:"profile_pic_url"`
	PhoneNumber      *string   `db:"phone_number" json:"phone_number"`
	DefaultAddressID *int      `db:"default_address_id" json:"default_address_id"`
	Tags             *string   `db:"tags" json:"tags"`
	Note             *string   `db:"note" json:"note"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
