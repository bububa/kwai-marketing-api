package dmp

// PopulationAccountsPushResponse 人群包跨账户推送 APIResponse
type PopulationAccountsPushResponse struct {
	// Success 推送成功的account_ids
	Success []uint64 `json:"success,omitempty"`
	// Fail 推送失败的account_ids
	Fail []uint64 `json:"fail,omitempty"`
}
