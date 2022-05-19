package tool

import (
	"net/url"
	"strconv"
)

// CreativeWordListRequest 获取可选的动态词包 API Request
type CreativeWordListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r CreativeWordListRequest) Url() string {
	return "v1/tool/creative_word/list"
}

// Encode implement GetRequest interface
func (r CreativeWordListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}
