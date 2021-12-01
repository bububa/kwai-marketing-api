package advertiser

import (
	"net/url"
	"strconv"
)

// FundGetRequest 获取广告主账户余额APIRequest
type FundGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r FundGetRequest) Url() string {
	return "v1/advertiser/fund/get"
}

// Encode implement GetRequest interface
func (r FundGetRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	return values.Encode()
}
