package dpa

// PageInfoSneak 分页信息
type PageInfoSneak struct {
	// PageSize 页大小
	PageSize int `json:"page_size,omitempty"`
	// CurrentPage 当前页码
	CurrentPage int `json:"current_page,omitempty"`
}
