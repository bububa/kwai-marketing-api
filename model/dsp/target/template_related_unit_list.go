package target

import "github.com/bububa/kwai-marketing-api/model"

// TemplateRelatedUnitListRequest 查询模板关联的广告列表接口 API Request
type TemplateRelatedUnitListRequest struct {
	// CurrentPage 当前页码
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 每页大小
	PageSize int `json:"page_size,omitempty"`
	// TaskID 任务ID
	TaskID uint64 `json:"task_id,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateRelatedUnitListRequest) Url() string {
	return "gw/dsp/target/template/related_unit_list"
}

// Encode implement PostRequest interface
func (r TemplateRelatedUnitListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateRelatedUnitListResponse 查询模板关联的广告列表接口 API Response
type TemplateRelatedUnitListResponse struct {
	// Details 模板同步列表
	Details []TemplateSyncUnit `json:"details,omitempty"`
	// CurrentPage 当前页码
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 每页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 总条数
	TotalCount int `json:"total_count,omitempty"`
}
