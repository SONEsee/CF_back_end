package dbschema

import "time"

type ProductVariantDBSchema struct {
	ID          int       `db:"id" json:"id"`
	ProductID   int       `db:"product_id" json:"product_id"`
	VariantName string    `db:"variant_name" json:"variant_name"`
	SkuCode     string    `db:"sku_code" json:"sku_code"`
	CfCode      string    `db:"cf_code" json:"cf_code"`
	Barcode     string    `db:"barcode" json:"barcode"`
	Price       float64   `db:"price" json:"price"`
	CostPrice   float64   `db:"cost_price" json:"cost_price"`
	WeightGrams int       `db:"weight_grams" json:"weight_grams"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
