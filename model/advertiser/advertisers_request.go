package advertiser

import (
	"net/url"
	"strconv"
)

// FundDailyFlowsRequest 广告主账号流水信息APIRequest
type AdvertisersRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r AdvertisersRequest) Url() string {
	return "gw/uc/v1/advertisers"
}

// Encode implement GetRequest interface
func (r AdvertisersRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))

	return values.Encode()
}
