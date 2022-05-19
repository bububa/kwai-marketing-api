package target

// Template 定向模板
type Template struct {
	// TemplateID 定向模板 ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateName 定向模板名称
	TemplateName string `json:"template_name,omitempty"`
	// CreateTime 定向模板创建时间，"2019-06-11 15:17:25"
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 定向模板最近修改时间
	UpdateTime string `json:"update_time,omitempty"`
	// AutoTarget 过滤定向模版; 默认：返回所有类型模版；false：只获取自定义定向模版；true ：只获取智能定向模版
	AutoTarget bool `json:"auto_target,omitempty"`
	// Target 定向信息
	Target *Target `json:"target,omitempty"`
	// UnitCount 绑定unit个数
	UnitCount int `json:"unit_count,omitempty"`
	// BehaviorInterest 行为兴趣定向
	BehaviorInterest *BehaviorInterest `json:"behavior_interest,omitempty"`
}
