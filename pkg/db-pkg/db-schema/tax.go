package dbschema

import "time"

type Tax struct {
	ID        int       `db:"id" json:"id"`
	NameTax   string    `db:"name_tax" json:"name_tax"`
	ValueTax  int       `db:"value_tax" json:"value_tax"`
	TaxDetail string    `db:"tax_detail" json:"tax_detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
