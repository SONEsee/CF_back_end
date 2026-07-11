package requestbody

type ShopRequestBody struct {
	ShopName    string `json:"shop_name" validate:"required,min=2,max=150"`
	OwnerUserID *int   `json:"owner_user_id" validate:"omitempty,gt=0"`
	Phone       string `json:"phone" validate:"omitempty,max=20"`
	Timezone    string `json:"timezone" validate:"omitempty,max=50"`
}

type ShopPatchRequest struct {
	ShopName    *string `json:"shop_name,omitempty" validate:"omitempty,min=2,max=150"`
	OwnerUserID *int    `json:"owner_user_id,omitempty" validate:"omitempty,gt=0"`
	Phone       *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Timezone    *string `json:"timezone,omitempty" validate:"omitempty,max=50"`
}

// ShopStatusRequest ໃຊ້ປ່ຽນສະຖານະຮ້ານ (ACTIVE / SUSPENDED / TRIAL) — ໃຊ້ແທນການລົບ ເນື່ອງຈາກ shops ບໍ່ມີ deleted_at
type ShopStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=ACTIVE SUSPENDED TRIAL"`
}
