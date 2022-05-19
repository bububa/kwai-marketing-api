package unit

import "encoding/json"

// UpdateDayBudgetRequest 修改广告组预算
type UpdateDayBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// DayBudget 单日预算金额; 广告组单日预算金额，单位：厘，指定0表示预算不限，默认为0；不小于100元，不超过100000000元，仅支持输入数字；修改预算不得低于该广告组当日花费的120%
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget不能同时传，均不能低于该计划下任一广告组出价
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateDayBudgetRequest) Url() string {
	return "v1/ad_unit/update/day_budget"
}

// Encode implement PostRequest interface
func (r UpdateDayBudgetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
