package page

// ListResponse 获取魔力建站落地页信息列表 API Response
type ListResponse struct {
	// TotalCount 总共条数
	TotalCount int `json:"total_count,omitempty"`
	// Details json array
	Details []Page `json:"details,omitempty"`
}
