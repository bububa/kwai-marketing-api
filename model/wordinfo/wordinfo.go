package wordinfo

// WordInfo 关键词
type WordInfo struct {
	// WordInfoID 关键词ID
	WordInfoID uint64 `json:"word_info_id,omitempty"`
	// Word 关键词
	Word string `json:"word,omitempty"`
	// MatchType 匹配类型
	// 1 - 精确匹配，2 - 短语匹配，3 - 广泛匹配
	MatchType int `json:"match_type,omitempty"`
	// ReviewStatus 审核状态
	// 1 - 审核中，2 - 审核通过，3 - 审核未通过，7 - 待送审
	ReviewStatus int `json:"review_status,omitempty"`
	// PutStatus 投放状态
	// 1 - 投放中，2 - 已暂停，3 - 已删除
	PutStatus int `json:"put_status,omitempty"`
	// Status 综合状态
	// 101 - 已删除，102 - 审核失败，103 - 审核中，104 - 已暂停，105 - 投放中，106 - 待送审
	Status int `json:"status,omitempty"`
	// ErrorReason 失败原因
	ErrorReason string `json:"error_reason,omitempty"`
}
