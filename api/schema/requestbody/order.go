package requestbody

// OrderItemInput — ລາຍການສິນຄ້າໃນ order, ບໍ່ມີ price ຍ້ອນ server ຈະ snapshot ລາຄາ product_variant ປັດຈຸບັນເອງ
type OrderItemInput struct {
	ProductVariantID int `json:"product_variant_id" validate:"required,gt=0"`
	BuyQuantity      int `json:"buy_quantity" validate:"required,gt=0"`
}

// OrderRequestBody — order_number/current_status/ຍອດເງິນທັງໝົດ server ຄິດໄລ່ເອງ, ບໍ່ຮັບຈາກ client
type OrderRequestBody struct {
	ShopID        int              `json:"shop_id" validate:"required,gt=0"`
	CustomerID    int              `json:"customer_id" validate:"required,gt=0"`
	LiveSessionID *int             `json:"live_session_id" validate:"omitempty,gt=0"`
	DiscountID    *int             `json:"discount_id" validate:"omitempty,gt=0"`
	ShippingFee   float64          `json:"shipping_fee" validate:"omitempty,gte=0"`
	Note          string           `json:"note"`
	Items         []OrderItemInput `json:"items" validate:"required,min=1,dive"`
}
