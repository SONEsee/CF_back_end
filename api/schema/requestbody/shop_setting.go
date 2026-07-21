package requestbody

// VatRate ຂຶ້ນຢູ່ກັບ column shop_settings.vat_rate NUMERIC(5,2) — ສູງສຸດ 999.99,
// ຕ້ອງຈຳກັດຢູ່ນີ້ບໍ່ດັ່ງນັ້ນຄ່າເກີນຂອບເຂດຈະຜ່ານ validate ແລ້ວໄປລົ້ມທີ່ຊັ້ນ DB ແທນ (22003 numeric field overflow)
type ShopSettingRequestBody struct {
	ShopID        int     `json:"shop_id" validate:"required,gt=0"`
	Currency      string  `json:"currency" validate:"omitempty,len=3"`
	VatRate       float64 `json:"vat_rate" validate:"omitempty,gte=0,lte=999.99"`
	AutoReplyMsg  string  `json:"auto_reply_msg"`
	BusinessHours string  `json:"business_hours" validate:"omitempty,json"`
}

type ShopSettingPatchRequest struct {
	Currency      *string  `json:"currency,omitempty" validate:"omitempty,len=3"`
	VatRate       *float64 `json:"vat_rate,omitempty" validate:"omitempty,gte=0,lte=999.99"`
	AutoReplyMsg  *string  `json:"auto_reply_msg,omitempty"`
	BusinessHours *string  `json:"business_hours,omitempty" validate:"omitempty,json"`
}
