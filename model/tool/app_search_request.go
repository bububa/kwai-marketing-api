package tool

import (
	"net/url"
	"strconv"
)

// AppSearchRequest 获取可选的应用定向 API Request
type AppSearchRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// AppName 应用名称; 支持模糊匹配
	AppName string `json:"app_name,omitempty"`
}

// Url implement GetRequest interface
func (r AppSearchRequest) Url() string {
	return "v1/tool/app/search"
}

// Encode implement GetRequest interface
func (r AppSearchRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	values.Set("app_name", r.AppName)
	return values.Encode()
}
