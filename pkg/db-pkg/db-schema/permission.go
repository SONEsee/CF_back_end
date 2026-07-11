package dbschema

type PermissionDBSchema struct {
	ID        int  `db:"id" json:"id"`
	RoleID    int  `db:"role_id" json:"role_id"`
	SubmenuID int  `db:"submenu_id" json:"submenu_id"`
	CanView   bool `db:"can_view" json:"can_view"`
	CanCreate bool `db:"can_create" json:"can_create"`
	CanUpdate bool `db:"can_update" json:"can_update"`
	CanDelete bool `db:"can_delete" json:"can_delete"`
}
