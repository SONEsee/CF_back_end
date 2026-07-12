package requestbody

type StockReservationRequestBody struct {
	ProductVariantID int    `json:"product_variant_id" validate:"required,gt=0"`
	CustomerID       *int   `json:"customer_id" validate:"omitempty,gt=0"`
	OrderItemID      *int   `json:"order_item_id" validate:"omitempty,gt=0"`
	ReservedQty      int    `json:"reserved_qty" validate:"required,gt=0"`
	ExpiresAt        string `json:"expires_at" validate:"required"`
}

// StockReservationStatusRequest ໃຊ້ຢືນຢັນຂາຍ (COMPLETED) ຫຼືປົດການຈອງ (EXPIRED) — ບໍ່ອະນຸຍາດປ່ຽນກັບໄປ HOLDING ດ້ວຍມື
type StockReservationStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=COMPLETED EXPIRED"`
}
