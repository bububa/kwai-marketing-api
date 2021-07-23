package report

import "encoding/json"

// AgentReportRequest 代理商数据APIRequest
type AgentReportRequest struct {
	// AgentID 代理商ID（注：非账户快手ID），在获取accessToken时返回
	AgentID int64 `json:"agent_id,omitempty"`
	// StartDate  过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	StartDate string `json:"start_date,omitempty"`
	// EndDate 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	EndDate string `json:"end_date,omitempty"`
	// Page 请求的页码，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认为20，最大支持1000
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest inferface
func (r AgentReportRequest) Url() string {
	return "v1/agent/report"
}

// Encode implement PostRequest interface
func (r AgentReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
