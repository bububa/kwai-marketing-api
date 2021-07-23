package report

// ReportResponse 数据报表APIResponse公用数据
type ReportResponse struct {
	// TotalCount 数据的总行数
	TotalCount int `json:"total_count,omitempty"`
	// Details 数据明细信息
	Details []Stat `json:"details,omitempty"`
}
