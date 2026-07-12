package dbschema

type ChatTemplateDBSchema struct {
	ID             int     `db:"id" json:"id"`
	ShopID         int     `db:"shop_id" json:"shop_id"`
	TriggerKeyword *string `db:"trigger_keyword" json:"trigger_keyword"`
	ResponseBody   string  `db:"response_body" json:"response_body"`
	IsActive       bool    `db:"is_active" json:"is_active"`
}
