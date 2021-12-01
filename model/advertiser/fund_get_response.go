package advertiser

// FundGetResponse 获取广告主账户余额APIResponse
type FundGetResponse struct {
	// Balance 账户总余额;单位：元
	Balance float64 `json:"balance,omitempty"`
}
