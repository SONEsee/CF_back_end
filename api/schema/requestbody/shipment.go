package requestbody

type ShipmentRequestBody struct {
	OrderID        int    `json:"order_id" validate:"required,gt=0"`
	CourierName    string `json:"courier_name" validate:"omitempty,max=50"`
	TrackingNumber string `json:"tracking_number" validate:"omitempty,max=100"`
	LabelPdfURL    string `json:"label_pdf_url" validate:"omitempty,url"`
}

type ShipmentPatchRequest struct {
	CourierName    *string `json:"courier_name,omitempty" validate:"omitempty,max=50"`
	TrackingNumber *string `json:"tracking_number,omitempty" validate:"omitempty,max=100"`
	LabelPdfURL    *string `json:"label_pdf_url,omitempty" validate:"omitempty,url"`
}

// ShipmentStatusRequest — ປ່ຽນ shipping_status: PICKED_UP ຈະ set shipped_at, DELIVERED ຈະ set delivered_at ອັດຕະໂນມັດ
type ShipmentStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=PICKED_UP DELIVERED"`
}
