package report

import "encoding/json"

// AudienceReportRequest 人群分析数据APIRequest
type AudienceReportRequest struct {
	ReportRequest
	// CampaignIDs 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// UnitIDs 广告组ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AudienceReportRequest) Url() string {
	return "v1/report/audience_report"
}

// Encode implement PostRequest interface
func (r AudienceReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
