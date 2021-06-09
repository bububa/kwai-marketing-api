package report

type MerchantDetailReportResponse struct {
	TotalCount int           `json:"total_count,omitempty"` // 数据的总行数
	Details    []MerchatStat `json:"details,omitempty"`     // 数据明细信息
}
