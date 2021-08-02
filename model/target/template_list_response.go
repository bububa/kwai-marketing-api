package target

// TemplateListResponse 查询定向模板接口 API Response
type TemplateListResponse struct {
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 列表
	Details []Template `json:"details,omitempty"`
}
