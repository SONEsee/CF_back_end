package requestbody

// CustomerAddressRequestBody — ບໍ່ມີ field is_default, ຕັ້ງໄດ້ຜ່ານ endpoint /customer/:id/default-address ເທົ່ານັ້ນ
type CustomerAddressRequestBody struct {
	CustomerID    int    `json:"customer_id" validate:"required,gt=0"`
	RecipientName string `json:"recipient_name" validate:"omitempty,max=150"`
	Phone         string `json:"phone" validate:"omitempty,max=20"`
	Address       string `json:"address"`
	SubDistrict   string `json:"sub_district" validate:"omitempty,max=100"`
	District      string `json:"district" validate:"omitempty,max=100"`
	Province      string `json:"province" validate:"omitempty,max=100"`
	PostalCode    string `json:"postal_code" validate:"omitempty,max=10"`
}

type CustomerAddressPatchRequest struct {
	RecipientName *string `json:"recipient_name,omitempty" validate:"omitempty,max=150"`
	Phone         *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Address       *string `json:"address,omitempty"`
	SubDistrict   *string `json:"sub_district,omitempty" validate:"omitempty,max=100"`
	District      *string `json:"district,omitempty" validate:"omitempty,max=100"`
	Province      *string `json:"province,omitempty" validate:"omitempty,max=100"`
	PostalCode    *string `json:"postal_code,omitempty" validate:"omitempty,max=10"`
}
