package tool

import (
	"net/url"
	"strconv"
)

// ConvertListRequest 获取可用的转化目标 API Request
type ConvertListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 转化目标类型; 1：JS布玛 2：Xpath 3：应用-SDK 7：应用-API
	Type int `json:"type,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r ConvertListRequest) Url() string {
	return "v1/tool/convert/list"
}

// Encode implement GetRequest interface
func (r ConvertListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("type", strconv.Itoa(r.Type))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
