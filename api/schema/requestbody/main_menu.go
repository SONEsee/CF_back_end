package requestbody

type MainMenuRequesBody struct {
	ModuleID int    `json:"module_id" validate:"required,gt=0"`
	NameMenu string `json:"menu_name" validate:"required,min=2,max=100"`
	IconMenu string `json:"icon_class" validate:"omitempty,max=50"`
}

type MainMenuPatchRequest struct {
	ModuleID *int    `json:"module_id,omitempty" validate:"omitempty,gt=0"`
	NameMenu *string `json:"menu_name,omitempty" validate:"omitempty,min=2,max=100"`
	IconMenu *string `json:"icon_class,omitempty" validate:"omitempty,max=50"`
}
