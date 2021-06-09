package advertiser

import "net/url"

type InfoRequest struct {
	AdvertiserID int64 `json:"advertiser_id,omitempty"` // 广告主ID
}

func (r InfoRequest) Url() string {
	return "v1/advertiser/info"
}

func (r InfoRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", r.AdvertiserID)
	return values.Encode()
}
