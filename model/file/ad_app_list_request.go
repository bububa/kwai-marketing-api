package file

import (
	"net/url"
	"strconv"
)

// AdAppListRequest 获取应用列表 API Request
type AdAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdAppListRequest) Url() string {
	return "v2/file/ad/app/list"
}

// Encode implement GetRequest interface
func (r AdAppListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
