package file

import (
	"encoding/json"
)

// AdAppListRequest 获取应用列表 API Request
type AdAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id"`
	// AppIDS 应用ID
	AppIDS string `json:"app_ids,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdAppListRequest) Url() string {
	return "v2/file/ad/app/list"
}

// Encode implement PostRequest interface
func (r AdAppListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
