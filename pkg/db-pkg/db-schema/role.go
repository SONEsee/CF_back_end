package dbschema

import "time"

type Role struct {
	ID        int       `db:"id" json:"id"`
	RoleName  string    `db:"role_name" json:"role_name"`
	Detail    string    `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
