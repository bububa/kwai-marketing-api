package report

// MerchantDetailReportResponse 小店通转化数据APIResponse
type MerchantDetailReportResponse struct {
	// TotalCount 数据的总行数
	TotalCount int `json:"total_count,omitempty"`
	// Details 数据明细信息
	Details []MerchantStat `json:"details,omitempty"`
}
