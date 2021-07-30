package unit

// BehaviorInterest 行为兴趣定向
type BehaviorInterest struct {
	// Keyword 行为定向关键词
	Keyword []Keyword `json:"keyword,omitempty"`
	// Label 行为定向，类目词
	Label []string `json:"label,omitempty"`
	// TimeType 在多少天内发生行为的用户; 0:7天 1：15天 2：30天 3：90天4：180天
	TimeType int `json:"time_type,omitempty"`
	// StrengthType 行为强度; 0：不限 1：高强度
	StrengthType int `json:"strength_type,omitempty"`
	// SceneType 行为场景; 1：社区 2：APP 4：推广
	SceneType int `json:"scene_type,omitempty"`
}
