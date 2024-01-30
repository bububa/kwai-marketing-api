package advertiser

import (
	"encoding/json"
)

// FundGetRequest 获取广告主账户余额APIRequest
type FundGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement PostRequest interface
func (r FundGetRequest) Url() string {
	return "v1/advertiser/fund/get"
}

// Encode implement PostRequest interface
func (r FundGetRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
