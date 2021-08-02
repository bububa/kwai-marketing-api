package asset

// AdvCardListResponse 获取高级创意列表
type AdvCardListResponse struct {
	// TotalCount 卡片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 卡片列表
	Details []AdvCard `json:"details,omitempty"`
}
