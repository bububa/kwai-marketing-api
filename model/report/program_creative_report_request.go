package report

import "encoding/json"

type ProgramCreativeReportRequest struct {
	ReportRequest
	CampaignIDs []int64 `json:"campaign_ids,omitempty"` // 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs     []int64 `json:"unit_ids,omitempty"`     // 广告组ID集，过滤筛选条件，单次查询数量不超过5000
}

func (r ProgramCreativeReportRequest) Url() string {
	return "v1/report/program_creative_report"
}

func (r ProgramCreativeReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
