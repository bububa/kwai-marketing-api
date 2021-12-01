package file

import (
	"net/url"
	"strconv"
)

// AdImageListRequest 查询图片接口list接口API Request
type AdImageListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间; 与end_date同时传或同时不传；
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	EndDate string `json:"end_date,omitempty"`
	// Page 请求的页码数
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdImageListRequest) Url() string {
	return "v1/file/ad/image/list"
}

// Encode implement GetRequest interface
func (r AdImageListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	if r.StartDate != "" && r.EndDate != "" {
		values.Set("start_date", r.StartDate)
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
