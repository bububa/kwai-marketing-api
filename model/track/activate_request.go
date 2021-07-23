package track

import (
	"net/url"
	"strconv"
)

// ActivateRequest 转化回传APIRequest
type ActivateRequest struct {
	// Callback 接口一接收的__CALLBACK__替换后的http地址中的callback参数（参考下方示例中标红处）。
	Callback string `json:"callback,omitempty"`
	// EventType 事件类型，转化事件和参数值的对应关系见下表，建议对接前与快手运营同学沟通确认。
	EventType int `json:"event_type,omitempty"`
	// EventType 事件时间，13位毫秒级时间戳。若请求中携带event_type参数，则必须同时携带event_time参数。
	EventTime int64 `json:"event_time,omitempty"`
	// PurchaseAmount 金额。当event_type为{3,12,13,14,15}，必须同时回传消费金额参数;单位元，可保留两位小数。
	PurchaseAmount float64 `json:"purchase_amount,omitempty"`
	// UserTagsAge 用户年龄，可选参数，整数，仅限保险行业使用。
	UserTagsAge int `json:"user_tags_age,omitempty"`
	// UserTagsInsurance 社保情况，可选参数，整数，仅限保险行业使用，取值：0-无社保、1-有社保、2-其他。
	UserTagsInsurance *int `json:"user_tags_insurance,omitempty"`
	// WeightedPurchaseAmount 加权付费金额，可选参数，双精度浮点型，仅限保险行业使用。
	WeightedPurchaseAmount float64 `json:"weighted_purchase_amount,omitempty"`
	// ActionReason 行为原因，可选参数，整数，当event_type为{120,121}时强烈建议回传，枚举值见下方表格。
	ActionReason int `json:"action_reason,omitempty"`
	// KeyActionCategory 关键行为类型，可选参数，整数，当 event_type=143 时强烈建议回传，枚举值见下方表格。
	KeyActionCategory int `json:"key_action_category,omitempty"`
	// KeyActionThreshold 关键行为定义对应的次数/数量/等级/金额等信息，可选参数，数字，当 event_type=143 时强烈建议回传，枚举值见下方表格
	KeyActionThreshold int `json:"key_action_threshold,omitempty"`
	// IsDirectMatch 为提升投放效果，快手希望广告主可以将归因至其它渠道的转化数据也回传至快手，帮助优化深度转化模型训练，上报此类转化数据时，需要增加参数“&is_direct_match=false”，表示“该用户看过快手广告，但转化归因到了其它平台”，此类数据不会计入投放报表展现。
	IsDirectMatch *bool `json:"is_direct_match,omitempty"`
}

// Encode implement GetRequest interface
func (r ActivateRequest) Encode() string {
	values := url.Values{}
	values.Set("callback", r.Callback)
	values.Set("event_type", strconv.Itoa(r.EventType))
	values.Set("event_time", strconv.FormatInt(r.EventTime, 10))
	if r.PurchaseAmount > 0.0000000001 || r.EventType == 3 || r.EventType == 12 || r.EventType == 13 || r.EventType == 14 || r.EventType == 15 {
		values.Set("purchase_amount", strconv.FormatFloat(r.PurchaseAmount, 'f', 2, 64))
	}
	if r.UserTagsAge > 0 {
		values.Set("user_tags_age", strconv.Itoa(r.UserTagsAge))
	}
	if r.UserTagsInsurance != nil {
		values.Set("user_tags_insurance", strconv.Itoa(*r.UserTagsInsurance))
	}
	if r.WeightedPurchaseAmount > 0.0000000001 {
		values.Set("weighted_purchase_amount", strconv.FormatFloat(r.WeightedPurchaseAmount, 'f', 2, 64))
	}
	if r.ActionReason > 0 {
		values.Set("action_reason", strconv.Itoa(r.ActionReason))
	}
	if r.KeyActionCategory > 0 {
		values.Set("key_action_category", strconv.Itoa(r.KeyActionCategory))
	}
	if r.KeyActionThreshold > 0 {
		values.Set("key_action_threshold", strconv.Itoa(r.KeyActionThreshold))
	}
	if r.IsDirectMatch != nil {
		isDirectMatch := "false"
		if *r.IsDirectMatch {
			isDirectMatch = "true"
		}
		values.Set("is_direct_match", isDirectMatch)
	}
	return values.Encode()
}
