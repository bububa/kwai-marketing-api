package report

import "encoding/json"

type AudienceReportRequest struct {
	ReportRequest
	CampaignIDs []int64 `json:"campaign_ids,omitempty"` // 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs     []int64 `json:"unit_ids,omitempty"`     // 广告组ID集，过滤筛选条件，单次查询数量不超过5000
}

func (r AudienceReportRequest) Url() string {
	return "v1/report/audience_report"
}

func (r AudienceReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
