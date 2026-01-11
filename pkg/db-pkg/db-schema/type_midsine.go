package dbschema

import "time"

type TypeMedicineDBSchema struct {
	ID          int       `db:"id_type" json:"id"`
	NameType    string    `db:"name_type" json:"name_type"`
	Detail_Type string    `db:"detail_type" json:"detail_type"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
