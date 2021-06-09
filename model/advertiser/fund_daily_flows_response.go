package advertiser

type FundDailyFlowsResponse struct {
	TotalCount int             `json:"total_count,omitempty"`
	Details    []FundDailyFlow `json:"details,omitempty"`
}

type FundDailyFlow struct {
	Date                        string  `json:"date,omitempty"`                           // 日期
	DailyCharge                 float64 `json:"daily_charge,omitempty"`                   // 总花费
	RealCharged                 float64 `json:"real_charged,omitempty"`                   // 充值花费;广告主的现金消耗和返点消耗，单位：元
	ContractRebateRealCharged   float64 `json:"contract_rebate_real_charged,omitempty"`   // 框返花费;年框广告主的框架返点的消耗，单位：元
	DirectRebateRealCharged     float64 `json:"direct_rebate_real_charged,omitempty"`     // 激励花费;广告主激励账户中余额的消耗，单位：元
	DailyTransferIn             float64 `json:"daily_transfer_in,omitempty"`              // 转入
	DailyTransferOut            float64 `json:"daily_transfer_out,omitempty"`             // 转出
	Balance                     float64 `json:"balance,omitempty"`                        // 日终结余
	RealRecharged               float64 `json:"real_recharged,omitempty"`                 // 充值转入
	ContractRebateRealRecharged float64 `json:"contract_rebate_real_recharged,omitempty"` // 框返转入
	DirectRebateRealRecharged   float64 `json:"direct_rebate_real_recharged,omitempty"`   // 激励转入
}
