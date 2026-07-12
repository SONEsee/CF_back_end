package requestbody

type OrderStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=PAYMENT_PENDING_VERIFY PAID PACKING SHIPPED CANCELLED"`
	Note   string `json:"note"`
}
