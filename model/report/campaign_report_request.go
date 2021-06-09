package report

import "encoding/json"

type CampaignReportRequest struct {
	ReportRequest
	CampaignIDs  []int64 `json:"campaign_ids,omitempty"`  // 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	CampaignType int     `json:"campaign_type,omitempty"` // 计划类型，过滤筛选条件1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。
}

func (r CampaignReportRequest) Url() string {
	return "v1/report/campaign_report"
}

func (r CampaignReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
