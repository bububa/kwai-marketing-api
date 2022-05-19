package campaign

import "encoding/json"

// CreateRequest 创建广告计划 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignName 广告计划名称; 长度为1-100个字符，计划名称不能重复
	CampaignName string `json:"campaign_name,omitempty"`
	// Type 计划类型; 2：提升应用安装 3：获取电商下单 4：推广品牌活动 5：收集销售线索 7：提高应用活跃 9：商品库推广 13：小店商品推广 14：直播推广
	Type int `json:"type,omitempty"`
	// SubType 商品广告类型;5：SDPA 4:DPA 当type为9时，sub_type必为4、5两者之一。（目前只支持类型5：SDPA）
	SubType int `json:"sub_type,omitempty"`
	// DayBudget 单日预算金额;单位：厘，指定0表示预算不限，默认为0；不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget_schedule不能同时传，不能低于该计划下任一广告组出价
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget不能同时传，均不能低于该计划下任一广告组出价
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
}

// Url implement PostRequest interface
func (r CreateRequest) Url() string {
	return "v2/campaign/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
