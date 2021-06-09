package report

type AgentReportResponse struct {
	TotalCount int         `json:"total_count,omitempty"` // 数据的总行数
	Details    []AgentStat `json:"details,omitempty"`     // 数据明细信息
}
