package advertiser

import (
	"encoding/json"
)

// InfoRequest 获取广告主信息APIRequest
type InfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement PostRequest interface
func (r InfoRequest) Url() string {
	return "v1/advertiser/info"
}

// Encode implement PostRequest interface
func (r InfoRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
