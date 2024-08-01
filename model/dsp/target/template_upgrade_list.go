package target

// TemplateUpgradeListRequest 查询待升级模板列表
type TemplateUpgradeListRequest struct{}

// Url implement PostRequest interface
func (r TemplateUpgradeListRequest) Url() string {
	return "gw/dsp/target/template/upgrade_list"
}

// Encode implement PostRequest interface
func (r TemplateUpgradeListRequest) Encode() []byte {
	return nil
}

// TemplateUpgradeItem 待升级模板
type TemplateUpgradeItem struct {
	// TargetChangeVO 模板升级详情
	TargetChangeVO *TargetChangeVO `json:"target_change_vo,omitempty"`
	// TemplateName 模板名称
	TemplateName string `json:"template_name,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// UnitCount 关联的广告组数量
	UnitCount int64 `json:"unit_count,omitempty"`
}

// TargetChangeVO 模板升级详情
type TargetChangeVO struct {
	// Celebrity 快手网红是否发生变更
	// true 表示快手网红之前有值被清空
	Celebrity bool `json:"celebrity,omitempty"`
	// BehaviorInterest 行为意向是否发生变更
	// true 表示行为意向之前有值被清空
	BehaviorInterest bool `json:"behavior_interest,omitempty"`
	// InteliExtendOption 智能定向是否发生变更
	// true 表示智能定向被打开
	InteliExtendOption bool `json:"inteli_extend_option,omitempty"`
}
