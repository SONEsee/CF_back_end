package dbschema

type SubMenuDBSchema struct {
	ID          int    `db:"id" json:"id"`
	MainMenuID  int    `db:"main_menu_id" json:"main_menu_id"`
	MainMenuName string  `db:"main_menu_name" json:"main_menu_name"`
	SubmenuName string `db:"submenu_name" json:"submenu_name"`
	RoutePath   string `db:"route_path" json:"route_path"`
}
