package requestbody

// StockMovementRequestBody — QtyChange ເປັນຄ່າມີເຄື່ອງໝາຍ (ບວກ=ເພີ່ມ, ລົບ=ຫຼຸດ) ຕໍ່ actual_qty.
// ບໍ່ຮັບ balance_after ຈາກ client — server ຄິດໄລ່ເອງ (ອ່ານ inventories ປັດຈຸບັນ + qty_change).
type StockMovementRequestBody struct {
	ProductVariantID int    `json:"product_variant_id" validate:"required,gt=0"`
	MovementType     string `json:"movement_type" validate:"required,oneof=IN OUT ADJUST RESERVE RELEASE"`
	QtyChange        int    `json:"qty_change" validate:"required,ne=0"`
	RefType          string `json:"ref_type" validate:"omitempty,max=30"`
	RefID            *int64 `json:"ref_id" validate:"omitempty,gt=0"`
	Note             string `json:"note" validate:"omitempty,max=255"`
}
