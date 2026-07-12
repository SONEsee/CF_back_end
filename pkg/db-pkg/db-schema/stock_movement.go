package dbschema

import "time"

type StockMovementDBSchema struct {
	ID               int64     `db:"id" json:"id"`
	ProductVariantID int       `db:"product_variant_id" json:"product_variant_id"`
	MovementType     string    `db:"movement_type" json:"movement_type"`
	QtyChange        int       `db:"qty_change" json:"qty_change"`
	BalanceAfter     int       `db:"balance_after" json:"balance_after"`
	RefType          string    `db:"ref_type" json:"ref_type"`
	RefID            *int64    `db:"ref_id" json:"ref_id"`
	Note             string    `db:"note" json:"note"`
	CreatedBy        *int      `db:"created_by" json:"created_by"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
