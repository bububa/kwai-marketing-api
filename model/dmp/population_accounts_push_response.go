package dmp

// PopulationAccountsPushResponse 人群包跨账户推送 APIResponse
type PopulationAccountsPushResponse struct {
	// Success 推送成功的account_ids
	Success []int64 `json:"success,omitempty"`
	// Fail 推送失败的account_ids
	Fail []int64 `json:"fail,omitempty"`
}
