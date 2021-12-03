package advertiser

import (
	"encoding/json"
)

// AdvertisersRequest 广告罗盘请求信息
type AdvertisersRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r AdvertisersRequest) Url() string {
	return "gw/uc/v1/advertisers"
}

// Encode implement GetRequest interface
func (r AdvertisersRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
