package unit

import "encoding/json"

// SuggestBidDetailRequest 获取广告组出价建议 API Request
type SuggestBidDetailRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SuggestBidParams 出价建议入参
	SuggestBidParams []SuggestBidParam `json:"suggest_bid_param,omitempty"`
}

// SuggestBidParam 出价建议入参
type SuggestBidParam struct {
	// BidType 1：CPM，2：CPC，6：OCPC(使用 OCPC 代表 OCPX)，10：OCPM，20：eCPC
	BidType int `json:"bid_type,omitempty"`
	// UnitID 新建时该值为0，编辑时该值为广告组id
	UnitID uint64 `json:"unit_id,omitempty"`
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// PredictCpaBid 保持和current_cpa_bid一致即可。
	PredictCpaBid int64 `json:"predict_cpa_bid,omitempty"`
	// CurrentCpaBid 新建广告组时，该值为0；编辑广告组时，该值为用户原来填写的浅度出价。单位厘
	CurrentCpaBid int64 `json:"current_cpa_bid,omitempty"`
	// OcpxActionType 0：未知，2：点击转化链接，10：曝光，11：点击，31：下载完成，53：提交线索，109：电话卡激活，137：量房，180：激活，190: 付费，191：首日 ROI，348：有效线索，383: 授信，384: 完件 715：微信复制;739:7 日付费次数
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
	// DeepConversionType 3: 付费，7: 次日留存，10: 完件, 11: 授信 ，0：无
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// PredictDeepConvBid 保持和curr_deep_bid一致即可
	PredictDeepConvBid int64 `json:"predict_deep_conv_bid,omitempty"`
	// CurrDeepBid 新建广告组时，该值为0；编辑广告组时，该值为用户原来填写的深度出价。单位厘
	CurrDeepBid int64 `json:"curr_deep_bid,omitempty"`
}

// Url implement PostRequest interface
func (r SuggestBidDetailRequest) Url() string {
	return "gw/dsp/v1/unit/suggestBid/detail"
}

// Encode implement PostRequest interface
func (r SuggestBidDetailRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// SuggestBidUnit 广告组出价建议
type SuggestBidUnit struct {
	// CanSuggest 是否展示出价建议
	CanSuggest bool `json:"can_suggest,omitempty"`
	// SuggestBid 浅度出价建议值
	SuggestBid int64 `json:"suggest_bid,omitempty"`
	// SuggestBidMin 建议区间最小值
	SuggestBidMin int64 `json:"suggest_bid_min,omitempty"`
	// SuggestBidMax 建议区间最大值
	SuggestBidMax int64 `json:"suggest_bid_max,omitempty"`
	// UnitID 广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// CanPredict 是否可以预测转化
	CanPredict bool `json:"can_predict,omitempty"`
	// SuggestDeepBid 深度出价建议值
	SuggestDeepBid int64 `json:"suggest_deep_bid,omitempty"`
	// BidPercentile 浅度出价分位值
	BidPercentile *BidPercentile `json:"bid_percentile,omitempty"`
	// DeepBidPercentile 深度出价分位值
	DeepBidPercentile *BidPercentile `json:"deep_bid_percentile,omitempty"`
}

// BidPercentile 出价分位值
type BidPercentile struct {
	// SuggestBidPencetile 建议出价所处分位值
	SuggestBidPercentile int64 `json:"suggest_bid_percentile,omitempty"`
	// Percentile 不同出价与分位值的对应
	Percentile []Percentile `json:"percentile,omitempty"`
}

// Percentile 出价与分位值的对应
type Percentile struct {
	// PercentileNum 出价对应分位值
	PerventileNum int `json:"percentile_num,omitempty"`
	// Bid 出价
	Bid int64 `json:"bid,omitempty"`
}
