package unit

import "encoding/json"

// UpdateBidRequest 修改广告组出价 API Request
type UpdateBidRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// Bid 出价; 广告组bid_type为CPC和eCPC时：不得低于0.2元，不得高于100元，单位：厘；广告组bid_type为OCPC时：行为出价不得低于1元；激活出价不得低于5元（白名单用户不得低于2元），单位：厘
	Bid int64 `json:"bid,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateBidRequest) Url() string {
	return "v1/ad_unit/update/bid"
}

// Encode implement PostRequest interface
func (r UpdateBidRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
