package asynctask

import "encoding/json"

// ListRequest 获取任务状态 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID（注：非账户快手ID），在获取accessToken时返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// 任务 ID 集，不超过 10 个
	TaskIDs []uint64 `json:"task_ids,omitempty"`
	// Page 请求的页码，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认为20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest inferface
func (r ListRequest) Url() string {
	return "v1/async_task/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// ListResponse 获取任务状态 API Response
type ListResponse struct {
	// TotalCount 任务总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 任务详情
	Details []Task `json:"details,omitempty"`
}
