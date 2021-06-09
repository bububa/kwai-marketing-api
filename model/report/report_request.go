package report

import "github.com/bububa/kwai-marketing-api/model"

type ReportRequest struct {
	AdvertiserID        int64                         `json:"advertiser_id,omitempty"`        // 广告主ID（注：非账户快手ID），在获取accessToken时返回
	StartDateMin        string                        `json:"start_date_min,omitempty"`       // 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	EndDateMin          string                        `json:"end_date_min,omitempty"`         // 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	StartDate           string                        `json:"start_date,omitempty"`           // 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	EndDate             string                        `json:"end_date,omitempty"`             // 过滤筛选条件，格式yyyy-MM-dd 可选时间范围参见文档上方说明
	TemporalGranularity model.TemporalGranularityType `json:"temporal_granularity,omitempty"` // 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据
	ReportDims          []string                      `json:"report_dims,omitempty"`          // "adScene"：按广告场景；不传/传空/传空数组：不限
	Page                int                           `json:"page,omitempty"`                 // 请求的页码，默认为 1
	PageSize            int                           `json:"page_size,omitempty"`            // 每页行数，默认为20，最大支持2000
}
