package advertiser

import (
	"fmt"
	"net/url"
)

type FundGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id,omitempty"` // 广告主ID
}

func (r FundGetRequest) Url() string {
	return "v1/advertiser/fund/get"
}

func (r FundGetRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", fmt.Sprint(r.AdvertiserID))
	return values.Encode()
}
