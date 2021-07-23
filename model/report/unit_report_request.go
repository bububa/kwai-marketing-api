package report

import "encoding/json"

// UnitReportRequest 广告单元数据APIRequest
type UnitReportRequest struct {
	ReportRequest
	// CampaignIDs 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	CampaignIDs []int64 `json:"campaign_ids,omitempty"`
	// CampaignType 计划类型，过滤筛选条件1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。
	CampaignType int `json:"campaign_type,omitempty"`
	// UnitIDs 广告组ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs []int64 `json:"unit_ids,omitempty"`
}

// Url implement PostRequest interface
func (r UnitReportRequest) Url() string {
	return "v1/report/unit_report"
}

// Encode implement PostRequest interface
func (r UnitReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
