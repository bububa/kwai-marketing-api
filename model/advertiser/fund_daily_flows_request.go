package advertiser

import (
	"net/url"
	"strconv"
)

type FundDailyFlowsRequest struct {
	AdvertiserID int64  `json:"advertiser_id,omitempty"`
	StartDate    string `json:"start_date,omitempty"` // 开始日期
	EndDate      string `json:"end_date,omitempty"`   // 结束日期
	Page         int    `json:"page,omitempty"`       // 查询的页码数
	PageSize     int    `json:"page_size,omitempty"`  // 单页行数
}

func (r FundDailyFlowsRequest) Url() string {
	return "v1/advertiser/fund/daily_flows"
}

func (r FundDailyFlowsRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserId, 10))
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if r.Page > 0 {
		values.Set("page", r.Page)
	}
	if r.PageSize > 0 {
		values.Set("page_size", r.PageSize)
	}
	return values.Encode()
}
