package report

import "encoding/json"

// MaterialReportRequest 广告素材数据API Request
type MaterialReportRequest struct {
	ReportRequest
	// CampaignIDs 广告计划ID集，过滤筛选条件，单次查询数量不超过5000
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// CampaignType 计划类型，过滤筛选条件1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。
	CampaignType int `json:"campaign_type,omitempty"`
	// UnitIDs 广告组ID集，过滤筛选条件，单次查询数量不超过5000
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// CreativeIDs 广告创意ID集，过滤筛选条件，单次查询数量不超过5000
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// PhotoIDs 视频ID集，过滤筛选条件，单次查询数量不超过5000
	PhotoIDs []uint64 `json:"photo_ids,omitempty"`
	// CoverIDs 封面ID集，过滤筛选条件，单次查询数量不超过5000
	CoverIDs []uint64 `json:"cover_ids,omitempty"`
	// CreativeMaterialType 素材类型：1 - 竖版视频; 2 - 横版视频; 3 - 便利贴;5 -竖版图片; 6- 横版图片; 9-小图；10-组图
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	// ViewType 报表类型：5 - 视频报表 7 - 封面报表 8 - 便利贴报表;16-图片报表；
	ViewType int `json:"view_type,omitempty"`
}

// Url implement PostRequest interface
func (r MaterialReportRequest) Url() string {
	return "v1/report/material_report"
}

// Encode implement PostRequest interface
func (r MaterialReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// MaterialReportResponse 广告素材数据API Response
type MaterialReportResponse struct {
	Code      int    `json:"code,omitempty"`       // 返回码
	Message   string `json:"message,omitempty"`    // 返回信息
	RequestId string `json:"request_id,omitempty"` // 请求id
	ReportResponse
}

// IsError detect if the response is an error
func (r MaterialReportResponse) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r MaterialReportResponse) Error() string {
	return r.Message
}
