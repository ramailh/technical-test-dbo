package model

type (
	PaginationRequest struct {
		Search  string `json:"search"`
		Limit   int    `json:"limit"`
		Page    int    `json:"page"`
		Order   string `json:"order"`
		OrderBy string `json:"order_by"`
	}

	Meta struct {
		Total       int `json:"total"`
		Limit       int `json:"limit"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	}
)
