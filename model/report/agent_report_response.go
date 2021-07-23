package report

// AgentReportResponse 代理商数据APIResponse
type AgentReportResponse struct {
	// TotalCount 数据的总行数
	TotalCount int `json:"total_count,omitempty"`
	// Details 数据明细信息
	Details []AgentStat `json:"details,omitempty"`
}
