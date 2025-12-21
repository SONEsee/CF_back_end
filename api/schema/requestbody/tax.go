package requestbody

type TaxRequestBody struct {
	ID        int    `json:"id"`
	NameTax   string `json:"name_tax" validate:"required,min=2,max=100"`
	ValueTax  int    `json:"value_tax" validate:"required,min=0"`
	TaxDetail string `json:"tax_detail"`
}
