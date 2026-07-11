package requestbody

type ShopSubscriptionRequestBody struct {
	ShopID    int    `json:"shop_id" validate:"required,gt=0"`
	PlanID    int    `json:"plan_id" validate:"required,gt=0"`
	StartDate string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"end_date" validate:"omitempty,datetime=2006-01-02"`
}

type ShopSubscriptionPatchRequest struct {
	PlanID  *int    `json:"plan_id,omitempty" validate:"omitempty,gt=0"`
	EndDate *string `json:"end_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

// ShopSubscriptionStatusRequest ໃຊ້ຍົກເລີກ/ປ່ຽນສະຖານະ — ໃຊ້ແທນການລົບ ເນື່ອງຈາກ shop_subscriptions ບໍ່ມີ deleted_at
type ShopSubscriptionStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=ACTIVE EXPIRED CANCELLED"`
}
