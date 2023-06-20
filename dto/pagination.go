package dto

type Pagination struct {
	Total       int64 `json:"total"`
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
}
