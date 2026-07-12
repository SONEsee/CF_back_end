package dbschema

import "time"

type CommentIntentDBSchema struct {
	ID                      int64     `db:"id" json:"id"`
	CommentRawID            int64     `db:"comment_raw_id" json:"comment_raw_id"`
	CustomerID              *int      `db:"customer_id" json:"customer_id"`
	MatchedProductVariantID *int      `db:"matched_product_variant_id" json:"matched_product_variant_id"`
	ParsedQty               *int      `db:"parsed_qty" json:"parsed_qty"`
	IntentStatus            string    `db:"intent_status" json:"intent_status"`
	ProcessedAt             time.Time `db:"processed_at" json:"processed_at"`
}
