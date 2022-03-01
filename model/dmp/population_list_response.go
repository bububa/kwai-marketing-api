package dmp

// PopulationListResponse 人群列表查询接口 API Response
type PopulationListResponse struct {
	// TotalCount 图片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 详情
	Details []Population `json:"details,omitempty"`
}
