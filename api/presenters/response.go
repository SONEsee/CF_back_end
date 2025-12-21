package presenters

import (
	"time"
)

const (
	SUCCESS = 1
	FAIL    = 0
)

// ເພີ່ມ struct Response ນີ້
type Response struct {
	Timestamp string      `json:"timestamp"`
	Status    int         `json:"status"`
	Items     interface{} `json:"items,omitempty"`
	Message   string      `json:"message,omitempty"`
	Error     string      `json:"error,omitempty"`
}

/*
Return success response - (can be modify as you want)
*/
func ResponseSuccess(data interface{}) map[string]interface{} {
	t := time.Now()
	return map[string]interface{}{
		"timestamp": t.Format("2006-01-02-15-04-05"),
		"status":    SUCCESS,
		"items":     data,
		"error":     nil,
	}
}

/*
Return list data with pagination infos
*/
// NOTE: if not pagination infos is not required, pass -1 to currentPage, currentPageTotalItem, totalPage
func ResponseSuccessListData(data interface{}, currentPage, currentPageTotalItem, totalItems, totalPage int) map[string]interface{} {
	t := time.Now()
	return map[string]interface{}{
		"timestamp": t.Format("2006-01-02-15-04-05"),
		"status":    SUCCESS,
		"items": map[string]interface{}{
			"list_data": data,
			"pagination": map[string]interface{}{
				"current_page":            currentPage,
				"current_page_total_item": currentPageTotalItem,
				"total_page":              totalItems,
				"total_items":             totalPage,
			},
		},
		"error": nil,
	}
}

func ResponseError(message, error string) Response {
	t := time.Now()
	return Response{
		Timestamp: t.Format("2006-01-02-15-04-05"),
		Status:    FAIL,
		Message:   message,
		Error:     error,
	}
}

func ResponseSuccessWithData(message string, data interface{}) Response {
	t := time.Now()
	return Response{
		Timestamp: t.Format("2006-01-02-15-04-05"),
		Status:    SUCCESS,
		Message:   message,
		Items:     data,
	}
}
