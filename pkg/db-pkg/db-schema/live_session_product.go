package dbschema

type LiveSessionProductDBSchema struct {
	ID               int      `db:"id" json:"id"`
	LiveSessionID    int      `db:"live_session_id" json:"live_session_id"`
	ProductVariantID int      `db:"product_variant_id" json:"product_variant_id"`
	LivePrice        *float64 `db:"live_price" json:"live_price"`
	CfCodeOverride   *string  `db:"cf_code_override" json:"cf_code_override"`
}
