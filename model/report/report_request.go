package report

import "github.com/Shinku-Chen/kwai-marketing-api/model"

// ReportRequest 数据报表APIRequest公用数据
type ReportRequest struct {
	// AdvertiserID 广告主ID（注：非账户快手ID），在获取accessToken时返回
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// StartDateMin 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	StartDateMin string `json:"start_date_min,omitempty"`
	// EndDateMin 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	EndDateMin string `json:"end_date_min,omitempty"`
	// StartDate 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	StartDate string `json:"start_date,omitempty"`
	// EndDate 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	EndDate string `json:"end_date,omitempty"`
	// TemporalGranularity 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据
	TemporalGranularity model.TemporalGranularityType `json:"temporal_granularity,omitempty"`
	// ReportDims "adScene"：按广告场景；不传/传空/传空数组：不限
	ReportDims []string `json:"report_dims,omitempty"`
	// Page 请求的页码，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认为20，最大支持2000
	PageSize int `json:"page_size,omitempty"`
}
