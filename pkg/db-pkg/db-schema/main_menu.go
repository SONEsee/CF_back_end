package dbschema

import "time"

type MainMenuDGSchema struct {
	ID        int       `db:"id" json:"id"`
	NameMenu  string    `db:"mame_menu" json:"name_menu"`
	IconMenu  string    `db:"icon_menu" json:"icon_menu"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type MainMenuWhitSubMenuSchema struct {
	ID       int             `db:"id" json:"id"`
	NameMenu string          `db:"mame_menu" json:"name_menu"`
	IconMenu string          `db:"icon_menu" json:"icon_menu"`
	SubMenu  []SubMenuSchema `json:"sub_menus"`
}
