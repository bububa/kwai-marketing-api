package report

import "encoding/json"

// QueryInfoReportRequest 广告素材数据API Request
type QueryInfoReportRequest struct {
	ReportRequest
	// UnitID 广告组 ID，过滤筛选条件。不可同时筛选unit_id和word_info_ids
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitIDs 广告组 ID 集，过滤筛选条件，单次查询数量不超过 5000
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// WordInfoIDs 推广关键词ID集，过滤筛选条件，单次查询数量不超过 5000。推广关键词ID集可通过获取关键词列表接口获取
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
	// SelectedColumns 自定义列
	SelectedColumns []string `json:"selected_columns,omitempty"`
	// Query 搜索词，过滤筛选条件，单次查询数量不超过5000
	Query []string `json:"query,omitempty"`
	// ReportDims "adScene"：按广告场景；"placementType"：按广告范围，快手/联盟;不传/传空/传空数组：不限
	ReportDims []string `json:"report_dims,omitempty"`
}

// Url implement PostRequest interface
func (r QueryInfoReportRequest) Url() string {
	return "gw/dsp/v1/report/query_word_report"
}

// Encode implement PostRequest interface
func (r QueryInfoReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// QueryInfoReportResponse 广告素材数据API Response
type QueryInfoReportResponse struct {
	Code      int             `json:"code,omitempty"`       // 返回码
	Message   string          `json:"message,omitempty"`    // 返回信息
	RequestId string          `json:"request_id,omitempty"` // 请求id
	Data      QueryInfoReport `json:"data"`
}

// IsError detect if the response is an error
func (r QueryInfoReportResponse) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r QueryInfoReportResponse) Error() string {
	return r.Message
}

type QueryInfoStat struct {
	// Query 搜索词
	Query string `json:",omitempty"`
	// Word 关键词文本
	Word string `json:"word,omitempty"`
	//TriggerType 触发方式：1:精确匹配,2:短语匹配,3:广泛匹配
	TriggerType int `json:"trigger_type	,omitempty"`
	Stat
}

// QueryInfoReport 数据报表APIResponse公用数据
type QueryInfoReport struct {
	// TotalCount 数据的总行数
	TotalCount int `json:"total_count,omitempty"`
	// Details 数据明细信息
	Details []QueryInfoStat `json:"details,omitempty"`
}
