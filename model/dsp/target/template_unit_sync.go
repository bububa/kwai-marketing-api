package target

import "github.com/bububa/kwai-marketing-api/model"

// TemplateUnitSyncRequest 定向模板同步 API Request
type TemplateUnitSyncRequest struct {
	// SyncUnitIDs 同步的广告组ID列表
	SyncUnitIDs []uint64 `json:"sync_unit_ids,omitempty"`
	// SyncUnitType 1：同步定向模板关联的所有广告组，此时sync_unit_ids应该为空或者不传；2：同步到定向模板下指定的部分广告组，此时sync_unit_ids应该有值；3：同步到定向模板下的排除掉 sync_unit_ids 列表中的所有广告组。当 sync_unit_type = 2/3 时，对应的 sync_unit_ids 字段可从接口「/rest/openapi/gw/dsp/target/template/related_unit_list」获取相关联的广告组ID
	SyncUnitType int `json:"sync_unit_type,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateUnitSyncRequest) Url() string {
	return "gw/dsp/target/template/unit_sync"
}

// Encode implement PostRequest interface
func (r TemplateUnitSyncRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateUnitSyncResponse 定向模板同步 API Response
type TemplateUnitSyncResponse struct {
	// AsyncTask 是否是异步任务
	AsyncTask bool `json:"async_task,omitempty"`
	// HasTask 是否存在同步任务
	// 仅当模板同步时同步到的广告组为空时，该字段为 true
	HasTask bool `json:"has_task,omitempty"`
	// TaskID 同步的任务ID
	// 返回的 task_id 可根据接口 「/rest/openapi/gw/dsp/target/template/sync_history」查询相关的同步详情
	TaskID uint64 `json:"task_id,omitempty"`
	// TemplateID 同步的模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// UnitErrorCount 同步失败的广告组数量
	// 仅同步任务时会返回此字段
	UnitErrorCount int `json:"unit_error_count,omitempty"`
	// UnitSuccessCount 同步成功的广告组数据
	// 仅同步任务时会返回此字段
	UnitSuccessCount int `json:"unit_success_count,omitempty"`
}
