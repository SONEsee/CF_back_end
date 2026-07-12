package requestbody

// CustomerRequestBody — default_address_id ບໍ່ໃສ່ໄວ້ນຳ, ຕັ້ງໄດ້ຜ່ານ endpoint /customer/:id/default-address ເທົ່ານັ້ນ
type CustomerRequestBody struct {
	ShopID           int    `json:"shop_id" validate:"required,gt=0"`
	SocialPlatformID string `json:"social_platform_id" validate:"omitempty,max=100"`
	CustomerName     string `json:"customer_name" validate:"omitempty,max=150"`
	ProfilePicURL    string `json:"profile_pic_url" validate:"omitempty,url"`
	PhoneNumber      string `json:"phone_number" validate:"omitempty,max=20"`
	Tags             string `json:"tags" validate:"omitempty,max=255"`
	Note             string `json:"note"`
}

type CustomerPatchRequest struct {
	SocialPlatformID *string `json:"social_platform_id,omitempty" validate:"omitempty,max=100"`
	CustomerName     *string `json:"customer_name,omitempty" validate:"omitempty,max=150"`
	ProfilePicURL    *string `json:"profile_pic_url,omitempty" validate:"omitempty,url"`
	PhoneNumber      *string `json:"phone_number,omitempty" validate:"omitempty,max=20"`
	Tags             *string `json:"tags,omitempty" validate:"omitempty,max=255"`
	Note             *string `json:"note,omitempty"`
}
