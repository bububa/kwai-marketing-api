package file

import (
	"encoding/json"
)

// AdImageListRequest 查询图片接口list接口API Request
type AdImageListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间; 与end_date同时传或同时不传；
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	EndDate string `json:"end_date,omitempty"`
	// Page 请求的页码数
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdImageListRequest) Url() string {
	return "v1/file/ad/image/list"
}

// Encode implement GetRequest interface
func (r AdImageListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
