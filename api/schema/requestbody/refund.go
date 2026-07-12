package requestbody

type RefundRequestBody struct {
	OrderID      int     `json:"order_id" validate:"required,gt=0"`
	Reason       string  `json:"reason" validate:"omitempty,max=255"`
	RefundAmount float64 `json:"refund_amount" validate:"required,gt=0"`
}

// RefundStatusRequest — REQUESTED->APPROVED->DONE, ຫຼື REQUESTED->REJECTED
type RefundStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=APPROVED DONE REJECTED"`
}
