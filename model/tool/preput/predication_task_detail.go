package preput

import "encoding/json"

// PredicationTaskDetailsRequest 投前预估详情 API Request
type PredicationTaskDetailsRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PredicationID 投前预估任务id
	PredicationID uint64 `json:"predication_id,omitempty"`
	// OcpxActionType 优化目标，53表示表单提交，180表示激活
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
}

// Url implement PostRequest interface
func (r PredicationTaskDetailsRequest) Url() string {
	return "/gw/dsp/v1/pre_put/predication_task/details"
}

// Encode implement PostRequest interface
func (r PredicationTaskDetailsRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// AdPredicationTaskDetail 投前预估任务详情
type AdPredicationTaskDetail struct {
	// PredicationID 预估任务id
	PredicationID uint64 `json:"predication_id,omitempty"`
	// PredicationTime 预估时间,单位ms
	PredicationTime int64 `json:"predication_time,omitempty"`
	// PhotoURL 视频url链接
	PhotoURL string `json:"photo_url,omitempty"`
	// PhotoPredicationStatus 任务预估状态, 1表示预测中，2表示预测结果已出，3表示特征已过期
	PhotoPredicationStatus int `json:"photo_predication_status,omitempty"`
	// PhotoPushedStatus 任务同步到视频库状态，0表示未推送，1表示已经推送
	PhotoPushedStatus int `json:"photo_pushed_status,omitempty"`
	// Title 视频标题
	Title string `json:"title,omitempty"`
	// Duration 视频时长,单位是ms
	Duration int64 `json:"duration,omitempty"`
	// OcpxActionType 优化目标，53表示表单提交，180表示激活
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
	// ContentQuality 内容分（素材观感的结果）
	ContentQuality int `json:"content_quality,omitempty"`
	// PhotoID 视频id
	PhotoID uint64 `json:"photo_id,omitempty"`
	// ContentQualityPromot .
	ContentQualityPromot string `json:"content_quality_promot,omitempty"`
	// ContentDuplicatePromot.
	ContentDuplicatePromot string `json:"content_duplicate_promot,omitempty"`
	// ContentDuplicate 素材是否重复
	ContentDuplicate bool `json:"content_duplicate,omitempty"`
}
