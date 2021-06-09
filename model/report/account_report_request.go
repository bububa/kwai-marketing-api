package report

import "encoding/json"

type AccountReportRequest struct {
	ReportRequest
}

func (r AccountReportRequest) Url() string {
	return "v1/report/account_report"
}

func (r AccountReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
