package creative

type ListAdvancedCreativeResponse struct {
	// TotalCount 数据总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 返回值详情
	Details []AdvancedCreative `json:"details,omitempty"`
}
