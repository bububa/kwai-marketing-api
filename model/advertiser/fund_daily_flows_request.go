package advertiser

import (
	"net/url"
	"strconv"
)

// FundDailyFlowsRequest 广告主账号流水信息APIRequest
type FundDailyFlowsRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// StartDate 开始日期
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期
	EndDate string `json:"end_date,omitempty"`
	// Page 查询的页码数
	Page int `json:"page,omitempty"`
	// PageSize 单页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r FundDailyFlowsRequest) Url() string {
	return "v1/advertiser/fund/daily_flows"
}

// Encode implement GetRequest interface
func (r FundDailyFlowsRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
