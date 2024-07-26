package report

import "github.com/bububa/kwai-marketing-api/model"

// WordInfoReportRequest 关键词报表 API Request
type WordInfoReportRequest struct {
	ReportRequest
	// UnitID 广告组 ID，过滤筛选条件。不可同时筛选unit_id和word_info_ids
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitIDs 广告组 ID 集，过滤筛选条件，单次查询数量不超过 5000
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// WordInfoIDs 推广关键词ID集，过滤筛选条件，单次查询数量不超过 5000。推广关键词ID集可通过获取关键词列表接口获取
	WordInfoIDs []uint64 `json:"word_ids,omitempty"`
	// ExtendInfo 	扩展查询选项,输入extendSearch可以查询智能扩词数据
	ExtendInfo []string `json:"extend_info,omitempty"`
}

// Url implement PostRequest interface
func (r WordInfoReportRequest) Url() string {
	return "gw/dsp/v1/report/word_info_report"
}

// Encode implement PostRequest interface
func (r WordInfoReportRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
