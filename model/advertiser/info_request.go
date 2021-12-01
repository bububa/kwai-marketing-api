package advertiser

import (
	"net/url"
	"strconv"
)

// InfoRequest 获取广告主信息APIRequest
type InfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r InfoRequest) Url() string {
	return "v1/advertiser/info"
}

// Encode implement GetRequest interface
func (r InfoRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	return values.Encode()
}
