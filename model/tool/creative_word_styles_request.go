package tool

import (
	"net/url"
	"strconv"
)

// CreativeWordStylesRequest 获取可选的封面贴纸样式 API Request
type CreativeWordStylesRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r CreativeWordStylesRequest) Url() string {
	return "v1/tool/creative_word/styles"
}

// Encode implement GetRequest interface
func (r CreativeWordStylesRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}
