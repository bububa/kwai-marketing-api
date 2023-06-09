package target_v2

// TemplateDetailsResponse 查询定向模板接口 API Response
type TemplateDetailsResponse struct {
	// TotalCount 总数
	TotalCount  int        `json:"total_count,omitempty"`
	CurrentPage int        `json:"current_page"`
	PageSize    int        `json:"page_size"`
	Details     []Template `json:"details,omitempty"`
}

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
	// Target 定向信息
	Target *Target `json:"target,omitempty"`
	// UnitCount 绑定unit个数
	UnitCount int `json:"unit_count,omitempty"`
}
