package target

import "github.com/bububa/kwai-marketing-api/model"

// TemplateSyncHistoryRequest 模板同步失败查询 API Request
type TemplateSyncHistoryRequest struct {
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
func (r TemplateSyncHistoryRequest) Url() string {
	return "gw/dsp/target/template/sync_history"
}

// Encode implement PostRequest interface
func (r TemplateSyncHistoryRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateSyncHistoryResponse 模板同步失败查询 API Response
type TemplateSyncHistoryResponse struct {
	// Details 模板同步失败列表
	Details []TemplateSyncUnit `json:"details,omitempty"`
	// CurrentPage 当前页码
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 每页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 总条数
	TotalCount int `json:"total_count,omitempty"`
}

// TemplateSyncUnit 模板同步信息
type TemplateSyncUnit struct {
	// FailedMsg 同步失败详情信息
	FailedMsg string `json:"failed_msg,omitempty"`
	// UnitViewStatusReason 广告组状态描述
	UnitViewStatusReason string `json:"unit_view_status_reason,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// AutoAdjust 计划是否开启自动调控
	AutoAdjust int `json:"auto_adjust,omitempty"`
	// AutoBuild 计划是否开启自动基建
	AutoBuild int `json:"auto_build,omitempty"`
	// AutoManage 计划是否开启智能投放
	AutoManage int `json:"auto_manage,omitempty"`
	// UnitViewStatus 广告组状态
	// -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，，11：审核中，12：审核未通过，14：已结束，15：已暂停，17：组超预算，19：未达投放时间，20：有效，22：不在投放时段
	UnitViewStatus int `json:"unit_view_status,omitempty"`
}
