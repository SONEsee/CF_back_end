package dbschema

import "time"

type InventoryDBSchema struct {
	ID               int       `db:"id" json:"id"`
	ProductVariantID int       `db:"product_variant_id" json:"product_variant_id"`
	ActualQty        int       `db:"actual_qty" json:"actual_qty"`
	ReservedQty      int       `db:"reserved_qty" json:"reserved_qty"`
	AvailableQty     int       `db:"available_qty" json:"available_qty"`
	ReorderLevel     int       `db:"reorder_level" json:"reorder_level"`
	LastUpdated      time.Time `db:"last_updated" json:"last_updated"`
}
