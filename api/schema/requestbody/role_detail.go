package requestbody

type RoleDetail struct {
	ID        int    `json:"id"`
	Sale      string `json:"sale" validate:"required,numeric,min=1,max=5"`
	New       string `json:"new" validate:"required,numeric,min=1,max=5"`
	Edit      bool   `json:"edit"`
	Delele    string `json:"delete" validate:"required,numeric,min=1,max=5"`
	Detail    bool   `json:"detail"`
	SubMenuID string `json:"submenu_id" validate:"required,numeric"`
	RoleID    string `json:"role_id" validate:"required,numeric"`
}
