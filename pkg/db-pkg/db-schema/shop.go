package dbschema

import "time"

type ShopDBSchema struct {
	ID          int       `db:"id" json:"id"`
	ShopName    string    `db:"shop_name" json:"shop_name"`
	OwnerUserID *int      `db:"owner_user_id" json:"owner_user_id"`
	Phone       string    `db:"phone" json:"phone"`
	Status      string    `db:"status" json:"status"`
	Timezone    string    `db:"timezone" json:"timezone"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
