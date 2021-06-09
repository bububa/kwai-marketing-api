package report

import "encoding/json"

type MaterialReportRequest struct {
	ReportRequest
	CampaignIDs          []int64 `json:"campaign_ids,omitempty"`           // 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	CampaignType         int     `json:"campaign_type,omitempty"`          // 计划类型，过滤筛选条件1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。
	UnitIDs              []int64 `json:"unit_ids,omitempty"`               // 广告组ID集，过滤筛选条件，单次查询数量不超过5000
	CreativeIDs          []int64 `json:"creative_ids,omitempty"`           // 广告创意ID集，过滤筛选条件，单次查询数量不超过5000
	PhotoIDs             []int64 `json:"photo_ids,omitempty"`              // 视频ID集，过滤筛选条件，单次查询数量不超过5000
	CoverIDs             []int64 `json:"cover_ids,omitempty"`              // 封面ID集，过滤筛选条件，单次查询数量不超过5000
	CreativeMaterialType int     `json:"creative_material_type,omitempty"` // 素材类型：1 - 竖版视频; 2 - 横版视频; 3 - 便利贴;5 -竖版图片; 6- 横版图片; 9-小图；10-组图
	ViewType             int     `json:"view_type,omitempty"`              // 报表类型：5 - 视频报表 7 - 封面报表 8 - 便利贴报表;16-图片报表；
}

func (r MaterialReportRequest) Url() string {
	return "v1/report/material_report"
}

func (r MaterialReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
