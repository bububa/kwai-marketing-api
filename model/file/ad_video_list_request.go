package file

import "encoding/json"

// AdVideoListRequest 查询视频接口list接口API Request
type AdVideoListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 id列表，不超过 100 个 id
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// NewStatus 视频状态s; 0：删除， -1：全部数据，包含删除 不传默认返回不含删除的数据
	NewStatus *int `json:"new_status,omitempty"`
	// StartDate 开始时间; 与end_date同时传或同时不传；
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	EndDate string `json:"end_date,omitempty"`
	// Page 请求的页码数
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoListRequest) Url() string {
	return "v1/file/ad/video/list"
}

// Encode implement PostRequest interface
func (r AdVideoListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
