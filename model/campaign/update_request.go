package campaign

import "encoding/json"

// UpdateRequest 修改广告计划 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告计划名称; 长度为1-100个字符，计划名称不能重复
	CampaignName string `json:"campaign_name,omitempty"`
	// DayBudget 单日预算金额;单位：厘，指定0表示预算不限，默认为0；不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget_schedule不能同时传，不能低于该计划下任一广告组出价
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget不能同时传，均不能低于该计划下任一广告组出价
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateRequest) Url() string {
	return "v2/campaign/update"
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
