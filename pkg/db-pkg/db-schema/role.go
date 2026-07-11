package dbschema

import "time"

type RoleDBSchema struct {
	ID          int       `db:"id" json:"id"`
	ShopID      *int      `db:"shop_id" json:"shop_id"`
	RoleName    string    `db:"role_name" json:"role_name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
