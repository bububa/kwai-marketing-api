package campaign

import "encoding/json"

// UpdateBudgetRequest 修改广告计划预算 API Request
type UpdateBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// DayBudget 单日预算金额;单位：厘，指定0表示预算不限，默认为0；不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget_schedule不能同时传，不能低于该计划下任一广告组出价
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget不能同时传，均不能低于该计划下任一广告组出价
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateBudgetRequest) Url() string {
	return "v1/campaign/update"
}

// Encode implement PostRequest interface
func (r UpdateBudgetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
