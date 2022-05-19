package report

import "encoding/json"

// ProgramCreativeReportRequest 程序化创意数据APIRequest
type ProgramCreativeReportRequest struct {
	ReportRequest
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"` // CampaignIDs 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs     []uint64 `json:"unit_ids,omitempty"`     // UnitIDs 广告组ID集，过滤筛选条件，单次查询数量不超过5000
}

// Url implement PostRequest interface
func (r ProgramCreativeReportRequest) Url() string {
	return "v1/report/program_creative_report"
}

// Encode implement PostRequest interface
func (r ProgramCreativeReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
