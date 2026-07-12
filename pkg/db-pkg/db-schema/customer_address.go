package dbschema

type CustomerAddressDBSchema struct {
	ID            int     `db:"id" json:"id"`
	CustomerID    int     `db:"customer_id" json:"customer_id"`
	RecipientName *string `db:"recipient_name" json:"recipient_name"`
	Phone         *string `db:"phone" json:"phone"`
	Address       *string `db:"address" json:"address"`
	SubDistrict   *string `db:"sub_district" json:"sub_district"`
	District      *string `db:"district" json:"district"`
	Province      *string `db:"province" json:"province"`
	PostalCode    *string `db:"postal_code" json:"postal_code"`
	IsDefault     bool    `db:"is_default" json:"is_default"`
}
