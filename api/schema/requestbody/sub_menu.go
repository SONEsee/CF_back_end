package requestbody

type SubMenuRequesBody struct {
	ID          int    `json:"id" `
	NameSubMenu string `json:"name_submenu" validate:"required"`
	URLSubMenu  string `json:"url_submenu" validate:"required,min=4,max=150"`
	IconSubMenu string `json:"icon_submenu" validate:"required,min=3,max=40"`
	MainMenuID  string `json:"main_menu_id" validate:"required,min=1,max=15"`
	Action      string `json:"action" validate:"omitempty,max=50"`
}

type SubMenuRequesBodyPact struct {
	NameSubMenu *string `json:"name_submenu,omitempty" validate:"omitempty,min=1,max=100"`
	URLSubMenu  *string `json:"url_submenu,omitempty" validate:"omitempty,min=1,max=255"`
	IconSubMenu *string `json:"icon_submenu,omitempty" validate:"omitempty,min=1,max=100"`
	MainMenuID  *string `json:"main_menu_id,omitempty" validate:"omitempty,min=1,max=10"`
	Action      *string `json:"action,omitempty" validate:"omitempty,min=1,max=20"`
}
