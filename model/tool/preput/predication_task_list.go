package preput

import "encoding/json"

// PredicationTaskListRequest 投前预估列表页接口 API Request
type PredicationTaskListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 分页参数
	Page struct {
		// CurrentPage 当前页码, 第一页是1
		CurrentPage int `json:"current_page,omitempty"`
		// PageSize 分页大小
		PageSize int `json:"page_size,omitempty"`
	} `json:"page,omitempty"`
	// TaskIDs 要筛选的任务id列表
	TaskIDs []uint64 `json:"task_ids,omitempty"`
	// TaskCreateTimeDuration 要筛选的创建时间范围，yyyy-MM-dd，不能超过近一周
	TaskCreateTimeDuration []string `json:"task_create_time_duration,omitempty"`
}

// Url implement PostRequest interface
func (r PredicationTaskListRequest) Url() string {
	return "/gw/dsp/v1/pre_put/predication_task/list"
}

// Encode implement PostRequest interface
func (r PredicationTaskListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// PredicationTaskListResponse 投前预估列表页接口 API Response
type PredicationTaskListResponse struct {
	// Tasks 投前预估列表页
	Tasks []AdPredicationTaskDetail `json:"tasks,omitempty"`
	// Page 列表页参数
	Page struct {
		// CurrentPage 当前页码, 第一页是1
		CurrentPage int `json:"current_page,omitempty"`
		// PageSize 分页大小
		PageSize int `json:"page_size,omitempty"`
		// TotalCount 数据总数
		TotalCount int `json:"total_count,omitempty"`
	} `json:"page,omitempty"`
}
