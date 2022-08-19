package unit

import "encoding/json"

// BidPredictRequest 获取广告组曝光、转化预估 API Request
type BidPredictRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	SuggestBidParam
}

// Url implement PostRequest interface
func (r BidPredictRequest) Url() string {
	return "/gw/dsp/v1/unit/bidPredict"
}

// Encode implement PostRequest interface
func (r BidPredictRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// BidPredict 广告组出价参考
type BidPredict struct {
	// CanPredict 是否可以预测转化
	CanPredict bool `json:"can_predict,omitempty"`
	// CanPredictConversion 是否可以预测出价-转化曲线
	CanPredictConversion bool `json:"can_predict_conversion,omitempty"`
	// CanPredictDelivery 是否可以预测出价-曝光曲线
	CanPredictDelivery bool `json:"can_predict_delivery,omitempty"`
	// PredictImpression 预估曝光数
	PredictImpression int64 `json:"predict_impression,omitempty"`
	// PredictConversion 预估转化数
	PredictConversion int64 `json:"predict_conversion,omitempty"`
	// SuggestBid 浅度出价建议值
	SuggestBid int64 `json:"suggest_bid,omitempty"`
	// SuggestBidMin 建议区间最小值
	SuggestBidMin int64 `json:"suggest_bid_min,omitempty"`
	// SuggestBidMax 建议区间最大值
	SuggestBidMax int64 `json:"suggest_bid_max,omitempty"`
	// TipsID 优化提示
	TipsID int64 `json:"tips_id,omitempty"`
}
