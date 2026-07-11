package dbschema

type MainMenuDGSchema struct {
	ID       int    `db:"id" json:"id"`
	ModuleID int    `db:"module_id" json:"module_id"`
	NameMenu string `db:"menu_name" json:"menu_name"`
	IconMenu string `db:"icon_class" json:"icon_class"`
}
