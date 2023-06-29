package target_v2

// BehaviorInterest 行为兴趣定向
type BehaviorInterest struct {
	// Behavior 行为定向
	Behavior *Behavior `json:"behavior,omitempty"`
	// Interest 兴趣定向
	Interest *Interest `json:"interest,omitempty"`
}

// Behavior 行为定向
type Behavior struct {
	// Keyword 行为定向关键词
	Keyword []Keyword `json:"keyword,omitempty"`
	// Label 行为定向，类目词
	Label []string `json:"label,omitempty"`
	// TimeType 在多少天内发生行为的用户; 0:7天 1：15天 2：30天 3：90天4：180天
	TimeType int `json:"time_type,omitempty"`
	// StrengthType 行为强度; 0：不限 1：高强度
	StrengthType int `json:"strength_type,omitempty"`
	// SceneType 行为场景; 1：社区 2：APP 4：推广
	SceneType []int `json:"scene_type,omitempty"`
}

// Interest 兴趣定向
type Interest struct {
	// Label 兴趣定向类目词; 根据/rest/openapi/v1/tool/label/behavior_interest接口获取。将兴趣类目id从最高层类目id开始，以“-”连接起来，假如有一个类目id为80202，父类目id为802，最高层类目id为8，则此时应该写"8-802-80202"；如果想全选最高层类目"8"底下的所有子类目，填"8"
	Label []string `json:"label,omitempty"`
	// StrengthType 兴趣标签强度; 0：不限 1：高强度
	StrengthType int `json:"strength_type,omitempty"`
}

// Keyword 行为定向关键词
type Keyword struct {
	// ID 关键词id
	ID uint64 `json:"id,omitempty"`
	// Name 关键词名称
	Name string `json:"name,omitempty"`
}
