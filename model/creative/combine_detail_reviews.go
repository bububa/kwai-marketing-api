package creative

// CombineDetailView 审核不通过和正在审核的创意组合
type CombineDetailView struct {
	// ID 创意id
	ID uint64 `json:"id,omitempty"`
	// PhotoID 视频id; 已加密
	PhotoID uint64 `json:"photo_id,omitempty"`
	// CoverUrl 封面url
	CoverUrl string `json:"cover_url,omitempty"`
	// Caption 作品广告语
	Caption string `json:"caption,omitempty"`
	// ReviewStatus 审核状态;1：审核中2：审核通过3：不通过
	ReviewStatus int `json:"review_status,omitempty"`
	// ReviewDetail 审核信息;里面是一个String类型数据，是审核信息
	ReviewDetail []string `json:"review_detail,omitempty"`
	// PutStatus 程序化创意审核状态; 程序化创意操作状态，1：投放，2：暂停，3：删除
	PutStatus int `json:"put_status,omitempty"`
}
