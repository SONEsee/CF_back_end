package dbschema

type ModuleDBSchema struct {
	ID           int    `db:"id" json:"id"`
	ModuleName   string `db:"module_name" json:"module_name"`
	DisplayOrder int    `db:"display_order" json:"display_order"`
}
