package requestbody

type InventoryPatchRequest struct {
	ReorderLevel *int `json:"reorder_level,omitempty" validate:"omitempty,gte=0"`
}
