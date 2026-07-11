package requestbody

type ModuleRequestBody struct {
	ModuleName   string `json:"module_name" validate:"required,min=2,max=100"`
	DisplayOrder int    `json:"display_order" validate:"omitempty,gte=0"`
}

type ModulePatchRequest struct {
	ModuleName   *string `json:"module_name,omitempty" validate:"omitempty,min=2,max=100"`
	DisplayOrder *int    `json:"display_order,omitempty" validate:"omitempty,gte=0"`
}
