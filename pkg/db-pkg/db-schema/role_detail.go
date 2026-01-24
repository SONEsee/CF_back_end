package dbschema

import "time"

type RoleDetail struct {
	ID        int    `db:"id" json:"id"`
	Sale      string `db:"sale" json:"sale"`
	New       string `db:"new" json:"new"`
	Edit      string `db:"edit" json:"edit"`
	Delele    string `db:"delete" json:"delete"`
	Detail    string `db:"detail" json:"detail"`
	SubMenuID int    `db:"submenu_id" json:"submenu_id"`

	RoleID    int       `db:"role_id" json:"role_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
