package tool

// ConvertListResponse 获取可用的转化目标 API Response
type ConvertListResponse struct {
	// TotalCount 结果总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 转化目标
	Details []Convert `json:"details,omitempty"`
}
