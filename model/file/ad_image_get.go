package file

import (
	"net/url"
	"strconv"
)

// AdImageGetRequest 查询图片信息get接口 API Request
type AdImageGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageToken 图片token
	ImageToken string `json:"image_token,omitempty"`
}

// Url implement GetRequest interface
func (r AdImageGetRequest) Url() string {
	return "v1/file/ad/image/get"
}

// Encode implement GetRequest interface
func (r AdImageGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("image_token", r.ImageToken)
	return values.Encode()
}
