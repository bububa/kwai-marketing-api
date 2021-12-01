package advertiser

import (
	"net/url"
	"strconv"
)

// WhiteListRequest 获取可选白名单接口 API Request
type WhiteListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Url implement PostRequest interface
func (r WhiteListRequest) Url() string {
	return "v1/advertiser/white_list"
}

// Encode implement PostRequest interface
func (r WhiteListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	return values.Encode()
}
