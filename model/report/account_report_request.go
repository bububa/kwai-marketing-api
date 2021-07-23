package report

import "encoding/json"

// AccountReportRequest 广告主数据APIRequest
type AccountReportRequest struct {
	ReportRequest
}

// Url implement PostRequest interface
func (r AccountReportRequest) Url() string {
	return "v1/report/account_report"
}

// Encode implement PostRequest interface
func (r AccountReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
