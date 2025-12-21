package requestbody

type TypeMedicine struct {
	ID         int    `json:"id_type"`
	NameType   string `json:"name_type" validate:"required,min=2,max=100"`
	DetailType string `json:"detail_type"`
}
