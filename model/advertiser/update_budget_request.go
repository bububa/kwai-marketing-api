package advertiser

import "encoding/json"

// UpdateBudgetRequest 修改账户预算APIRequest
type UpdateBudgetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// DayBudget 单日预算 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该账户当日花费的120%，与day_budget不能同时传
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule  单日预算金额; 广告组单日预算金额，单位：厘，指定0表示预算不限，默认为0；不小于100元，不超过100000000元，仅支持输入数字；修改预算不得低于该广告组当日花费的120%
	DayBudgetSchedule int64 `json:"day_budget_schedule,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateBudgetRequest) Url() string {
	return "v1/advertiser/update/budget"
}

// Encode implement PostRequest interface
func (r UpdateBudgetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
