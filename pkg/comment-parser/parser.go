package commentparser

import (
	"regexp"
	"strconv"
)

// cfPattern ຈັບຮູບແບບ "CF <code> x<qty>" — ບໍ່ສົນໃຈໂຕພິມນ້ອຍ/ໃຫຍ່, ບໍ່ສົນໃຈວັກລະຫວ່າງ (ຫຼືບໍ່ມີວັກ):
//
//	"CF ABC123 x2", "cf abc123x1", "CF ABC123 X 2" ລ້ວນ match
var cfPattern = regexp.MustCompile(`(?i)cf\s*([a-z0-9]+?)\s*x\s*(\d+)`)

// ParseCFComment ພະຍາຍາມດຶງລະຫັດສິນຄ້າ+ຈຳນວນອອກຈາກຂໍ້ຄວາມຄອມເມັນ.
// matched=false ໝາຍວ່າຂໍ້ຄວາມນີ້ບໍ່ແມ່ນຮູບແບບ CF (ຄອມເມັນທົ່ວໄປ) — caller ຄວນຖືວ່າເປັນ INVALID_CODE
func ParseCFComment(message string) (code string, qty int, matched bool) {
	m := cfPattern.FindStringSubmatch(message)
	if m == nil {
		return "", 0, false
	}
	parsedQty, err := strconv.Atoi(m[2])
	if err != nil || parsedQty <= 0 {
		return "", 0, false
	}
	return m[1], parsedQty, true
}
