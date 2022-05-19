package creative

// AdvancedProgramReviewDetail 获取程序化创意2.0审核信息
type AdvancedProgramReviewDetail struct {
	// UnitID 当前的程序化创意的广告组id
	UnitID uint64 `json:"unit_id,omitempty"`
	// Slogans 审核不通过的封面广告语
	Slogans []string `json:"slogans,omitempty"`
	// CombineDetailViews 审核不通过和正在审核的创意组合
	CombineDetailViews []CombineDetailView `json:"combine_detail_views,omitempty"`
}
