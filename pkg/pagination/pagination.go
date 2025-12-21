package pagination

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginationResult struct {
	CurrentPage          int `json:"current_page"`
	CurrentPageTotalItem int `json:"current_page_total_item"`
	TotalPage            int `json:"total_page"`
	TotalItems           int `json:"total_items"`
}

func NewPaginationParams(page, pageSize int) *PaginationParams {

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return &PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PaginationParams) GetLimit() int {
	return p.PageSize
}

func (p *PaginationParams) CalculatePagination(totalItems, currentPageItems int) *PaginationResult {
	totalPages := (totalItems + p.PageSize - 1) / p.PageSize

	return &PaginationResult{
		CurrentPage:          p.Page,
		CurrentPageTotalItem: currentPageItems,
		TotalPage:            totalPages,
		TotalItems:           totalItems,
	}
}

func (p *PaginationParams) IsValid() bool {
	return p.Page > 0 && p.PageSize > 0
}
