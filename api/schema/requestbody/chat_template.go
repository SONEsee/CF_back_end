package requestbody

type ChatTemplateRequestBody struct {
	ShopID         int    `json:"shop_id" validate:"required,gt=0"`
	TriggerKeyword string `json:"trigger_keyword" validate:"omitempty,max=100"`
	ResponseBody   string `json:"response_body" validate:"required"`
}

type ChatTemplatePatchRequest struct {
	TriggerKeyword *string `json:"trigger_keyword,omitempty" validate:"omitempty,max=100"`
	ResponseBody   *string `json:"response_body,omitempty"`
	IsActive       *bool   `json:"is_active,omitempty"`
}
