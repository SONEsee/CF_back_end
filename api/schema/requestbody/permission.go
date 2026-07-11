package requestbody

type PermissionRequestBody struct {
	RoleID    int  `json:"role_id" validate:"required,gt=0"`
	SubmenuID int  `json:"submenu_id" validate:"required,gt=0"`
	CanView   bool `json:"can_view"`
	CanCreate bool `json:"can_create"`
	CanUpdate bool `json:"can_update"`
	CanDelete bool `json:"can_delete"`
}

type PermissionPatchRequest struct {
	CanView   *bool `json:"can_view,omitempty"`
	CanCreate *bool `json:"can_create,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	CanDelete *bool `json:"can_delete,omitempty"`
}
