package asynctask

import "encoding/json"

// CreateRequest 创建异步任务 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID（注：非账户快手ID），在获取accessToken时返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskName 任务名称，最大不超过 50字符，不能为空；每个账户提交的任务名称不能重复
	TaskName string `json:"task_name,omitempty"`
	// TaskParams 任务参数
	TaskParams TaskParams `json:"task_params,omitempty"`
}

// Url implement PostRequest inferface
func (r CreateRequest) Url() string {
	return "v1/async_task/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// CreateResponse 创建异步任务 API Response
type CreateResponse struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`

	// TaskID 任务 ID
	TaskID uint64 `json:"task_id,omitempty"`
}
