package preput

import "encoding/json"

// PredicationTaskCreateRequest 创建投前预估任务 API Request
type PredicationTaskCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Tasks 投前预估创建任务的参数
	Tasks []AdPredicationTask `json:"tasks,omitempty"`
}

// Url implement PostRequest interface
func (r PredicationTaskCreateRequest) Url() string {
	return "/gw/dsp/v1/pre_put/predication_task/create"
}

// Encode implement PostRequest interface
func (r PredicationTaskCreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// RealTaskResult 投前预估创建结果
type RealTaskResult struct {
	// Successful 操作成功任务数目
	Successful int64 `json:"successful,omitempty"`
	// Failed 操作失败任务数目
	Failed int64 `json:"failed,omitempty"`
	// TaskFailedReasonList 操作任务成功失败详情列表
	TaskFailedReasonList []TaskFailedReason `json:"task_failed_reason_list,omitempty"`
}

// TaskFailedReason 操作任务成功失败详情
type TaskFailedReason struct {
	// Message 具体预估任务错误原因
	Message string `json:"message,omitempty"`
	// ReferTaskID md5重复的时候对应的任务id
	ReferTaskID uint64 `json:"refer_task_id,omitempty"`
	// PhotoPredicationStatus 重复任务的预测结果，1表示预测中，2表示预测结果已出，3表示特征已过期
	PhotoPredicationStatus int `json:"photo_predication_status,omitempty"`
	// OcpxActionType 重复任务的优化类型，53表示表单提交，180表示激活
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
}
