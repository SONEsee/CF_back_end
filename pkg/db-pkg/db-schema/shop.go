package dbschema

import "time"

type ShopDBSchema struct {
	ID          int       `db:"id" json:"id"`
	ShopName    string    `db:"shop_name" json:"shop_name"`
	OwnerUserID *int      `db:"owner_user_id" json:"owner_user_id"`
	Phone       string    `db:"phone" json:"phone"`
	Status      string    `db:"status" json:"status"`
	Timezone    string    `db:"timezone" json:"timezone"`
	ImageURL    *string   `db:"image_url" json:"image_url"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// ShopOptionDBSchema ເປັນ shape ເບົາໆສຳລັບ dropdown/autocomplete (ບໍ່ແມ່ນ list ຫຼັກ) — ບໍ່ມີ pagination
type ShopOptionDBSchema struct {
	ID       int    `db:"id" json:"id"`
	ShopName string `db:"shop_name" json:"shop_name"`
}
