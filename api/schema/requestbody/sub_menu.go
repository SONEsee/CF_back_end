package requestbody

type SubMenuRequesBody struct {
	MainMenuID  int    `json:"main_menu_id" validate:"required,gt=0"`
	SubmenuName string `json:"submenu_name" validate:"required,min=2,max=100"`
	RoutePath   string `json:"route_path" validate:"omitempty,max=255"`
}

type SubMenuPatchRequest struct {
	MainMenuID  *int    `json:"main_menu_id,omitempty" validate:"omitempty,gt=0"`
	SubmenuName *string `json:"submenu_name,omitempty" validate:"omitempty,min=2,max=100"`
	RoutePath   *string `json:"route_path,omitempty" validate:"omitempty,max=255"`
}
