package requestbody

// CommentIntentRequestBody — ຮັບຄ່າທີ່ຈັບຄູ່ໄວ້ແລ້ວ (ຈາກ staff ຫຼື parser ພາຍນອກໃນອະນາຄົດ), server ຮັບຜິດຊອບກວດ+ຈອງ stock:
//   - MatchedProductVariantID/ParsedQty ວ່າງ → intent_status=INVALID_CODE
//   - ມີແຕ່ stock ບໍ່ພຽງພໍ → intent_status=OUT_OF_STOCK (ບໍ່ຈອງ)
//   - ຈອງໄດ້ → intent_status=CF_SUCCESS (ຈອງ stock ຈິງ)
type CommentIntentRequestBody struct {
	CommentRawID            int64 `json:"comment_raw_id" validate:"required,gt=0"`
	CustomerID              *int  `json:"customer_id" validate:"omitempty,gt=0"`
	MatchedProductVariantID *int  `json:"matched_product_variant_id" validate:"omitempty,gt=0"`
	ParsedQty               *int  `json:"parsed_qty" validate:"omitempty,gt=0"`
}
