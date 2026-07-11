package dbschema

type ProductCategoryDBSchema struct {
	ID        int    `db:"id" json:"id"`
	ShopID    int    `db:"shop_id" json:"shop_id"`
	ParentID  *int   `db:"parent_id" json:"parent_id"`
	Name      string `db:"name" json:"name"`
	SortOrder int    `db:"sort_order" json:"sort_order"`
}
