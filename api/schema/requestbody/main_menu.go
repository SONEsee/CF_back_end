package requestbody

type MainMenuRequesBody struct {
	ID       int    `json:"id" `
	NameMenu string `json:"name_menu" validate:"required,min=2,max=100"`
	IconMenu string `json:"icon_menu" validate:"required,min=2,max=50"`
}

type MainMenuWhitSubMenuRequesBody struct {
	ID       int                 `json:"id" validate:"required"`
	NameMenu string              `json:"name_menu" validate:"required,min=2,max=100"`
	IconMenu string              `json:"icon_menu" validate:"required,min=2,max=50"`
	SubMenu  []SubMenuRequesBody `json:"sub_menu" validate:"omitempty, dive"`
}
