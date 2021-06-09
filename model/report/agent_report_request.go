package report

import "encoding/json"

type AgentReportRequest struct {
	AgentID   int64  `json:'agent_id,omitempty'`   // 代理商ID（注：非账户快手ID），在获取accessToken时返回
	StartDate string `json:"start_date,omitempty"` // 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	EndDate   string `json:"end_date,omitempty"`   // 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	Page      int    `json:"page,omitempty"`       // 请求的页码，默认为 1
	PageSize  int    `json:"page_size,omitempty"`  // 每页行数，默认为20，最大支持1000
}

func (r AgentReportRequest) Url() string {
	return "v1/agent/report"
}

func (r AgentReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
