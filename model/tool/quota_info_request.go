package tool

import (
	"encoding/json"
)

type QuotaInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r QuotaInfoRequest) Url() string {
	return "gw/dsp/quota/info"
}

// Encode implement GetRequest interface
// Encode implement PostRequest interface
func (r QuotaInfoRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
