package asset

// AdvCardCreateResponse 创建高级创意接口
type AdvCardCreateResponse struct {
	// AdvList 卡片 id 数组
	AdvList []uint64 `json:"adv_list,omitempty"`
}
