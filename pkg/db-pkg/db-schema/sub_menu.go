package dbschema

import "time"

type SubMenuSchema struct {
	ID          int       `db:"id" json:"id"`
	NameSubMenu string    `db:"name_submenu" json:"name_submenu"`
	IconSubMenu string    `db:"icon_submenu" json:"icon_submenu"`
	URLSubMenu  string    `db:"url_submenu" json:"url_submenu"`
	Action      string    `db:"action" json:"action"`
	MainMenuID  int       `db:"main_menu_id" json:"main_menu_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
