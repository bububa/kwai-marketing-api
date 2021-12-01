package advertiser

// FundDailyFlowsResponse 广告主账号流水信息APIResponse
type FundDailyFlowsResponse struct {
	TotalCount int             `json:"total_count,omitempty"`
	Details    []FundDailyFlow `json:"details,omitempty"`
}

// FundDailyFlow 广告主账号流水数据
type FundDailyFlow struct {
	// Date 日期
	Date string `json:"date,omitempty"`
	// DailyCharge 总花费
	DailyCharge float64 `json:"daily_charge,omitempty"`
	// RealCharged 充值花费;广告主的现金消耗和返点消耗，单位：元
	RealCharged float64 `json:"real_charged,omitempty"`
	// ContractRebateRealCharged 框返花费;年框广告主的框架返点的消耗，单位：元
	ContractRebateRealCharged float64 `json:"contract_rebate_real_charged,omitempty"`
	// DirectRebateRealCharged 激励花费;广告主激励账户中余额的消耗，单位：元
	DirectRebateRealCharged float64 `json:"direct_rebate_real_charged,omitempty"`
	// DailyTransferIn 转入
	DailyTransferIn float64 `json:"daily_transfer_in,omitempty"`
	// DailyTransferOut 转出
	DailyTransferOut float64 `json:"daily_transfer_out,omitempty"`
	// Balance 日终结余
	Balance       float64 `json:"balance,omitempty"`
	RealRecharged float64 `json:"real_recharged,omitempty"`
	// RealRecharged 充值转入
	// ContractRebateRealRecharged 框返转入
	ContractRebateRealRecharged float64 `json:"contract_rebate_real_recharged,omitempty"`
	// DirectRebateRealRecharged 激励转入
	DirectRebateRealRecharged float64 `json:"direct_rebate_real_recharged,omitempty"`
}
