package campaign

// ListResponse 获取广告计划信息 API Response
type ListResponse struct {
	// TotalCount 数据总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 返回值详情
	Details []Campaign `json:"details,omitempty"`
}
