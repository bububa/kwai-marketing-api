package target

import (
	"net/url"
	"strconv"
)

// TemplateListRequest 查询定向模板接口 API Request
type TemplateListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 请求的页码，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r TemplateListRequest) Url() string {
	return "v1/target/template/list"
}

// Encode implement GetRequest interface
func (r TemplateListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
