package creative

// AdvancedProgramListResponse 获取程序化创意2.0信息 API Response
type AdvancedProgramListResponse struct {
	// TotalCount 数据总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 返回值详情
	Details []AdvancedProgramCreative `json:"details,omitempty"`
}
