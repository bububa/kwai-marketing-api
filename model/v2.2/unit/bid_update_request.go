package unit

import "encoding/json"

type BidUpdateRequest struct {
	//advertiser_id	long	必填	广告主 ID	在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//unit_ids	long	必填	广告计划 ID
	UnitIds []uint64 `json:"unit_ids"`
	//广告组 bid_type 为 CPC 和 eCPC 时：不得低于 0.2 元，不得高于 100 元，单位：厘；广告组 bid_type 为 OCPC 时：行为出价不得低于 1 元；激活出价不得低于 5 元（白名单用户不得低于 2 元），单位：厘
	Bid uint64 `json:"bid"`
}

// Url implement PostRequest interface
func (r BidUpdateRequest) Url() string {
	return "v1/ad_unit/update/bid"
}

// Encode implement PostRequest interface
func (r BidUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
