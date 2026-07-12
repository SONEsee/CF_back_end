package requestbody

type SetDefaultAddressRequest struct {
	AddressID int64 `json:"address_id" validate:"required,gt=0"`
}
