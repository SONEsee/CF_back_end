package dbschema

type ProductImageDBSchema struct {
	ID        int    `db:"id" json:"id"`
	ProductID int    `db:"product_id" json:"product_id"`
	ImageURL  string `db:"image_url" json:"image_url"`
	SortOrder int    `db:"sort_order" json:"sort_order"`
}
