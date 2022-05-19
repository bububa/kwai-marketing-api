package region

import (
	"net/url"
	"strconv"
)

// DistrictListRequest 获取商圈地域定向 API Request
type DistrictListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r DistrictListRequest) Url() string {
	return "v1/region/district/list"
}

// Encode implement GetRequest interface
func (r DistrictListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}
