package dbschema

import "time"

type ProductDBSchema struct {
	ID           int       `db:"id" json:"id"`
	ShopID       int       `db:"shop_id" json:"shop_id"`
	CategoryID   *int      `db:"category_id" json:"category_id"`
	ProductName  string    `db:"product_name" json:"product_name"`
	Description  string    `db:"description" json:"description"`
	ImageMainURL string    `db:"image_main_url" json:"image_main_url"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
