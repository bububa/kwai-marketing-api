package report

import "encoding/json"

// WordInfoReportRequest 广告素材数据API Request
type WordInfoReportRequest struct {
	ReportRequest
	// UnitID 广告组 ID，过滤筛选条件。不可同时筛选unit_id和word_info_ids
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitIDs 广告组 ID 集，过滤筛选条件，单次查询数量不超过 5000
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// WordInfoIDs 推广关键词ID集，过滤筛选条件，单次查询数量不超过 5000。推广关键词ID集可通过获取关键词列表接口获取
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
	// ExtendInfo 扩展查询选项,输入extendSearch可以查询智能扩词数据
	ExtendInfo []string `json:"extend_info,omitempty"`
	// SelectedColumns 自定义列
	SelectedColumns []string `json:"selected_columns,omitempty"`
}

// Url implement PostRequest interface
func (r WordInfoReportRequest) Url() string {
	return "gw/dsp/v1/report/word_info_report"
}

// Encode implement PostRequest interface
func (r WordInfoReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// WordInfoReportResponse 广告素材数据API Response
type WordInfoReportResponse struct {
	Code      int    `json:"code,omitempty"`       // 返回码
	Message   string `json:"message,omitempty"`    // 返回信息
	RequestId string `json:"request_id,omitempty"` // 请求id
	Data      WordInfoReport
}

// IsError detect if the response is an error
func (r WordInfoReportResponse) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r WordInfoReportResponse) Error() string {
	return r.Message
}

type WordInfoStat struct {
	//WordInfoId 推广关键词ID
	WordInfoId uint64 `json:"word_info_id,omitempty"`
	//word 关键词文本
	Word string `json:"word,omitempty"`
	//MatchType 匹配方式：1:精确匹配,2:短语匹配,3:广泛匹配
	MatchType int `json:"match_type,omitempty"`
	Stat
}

// WordInfoReport 数据报表APIResponse公用数据
type WordInfoReport struct {
	// TotalCount 数据的总行数
	TotalCount int `json:"total_count,omitempty"`
	// Details 数据明细信息
	Details []WordInfoStat `json:"details,omitempty"`
}
