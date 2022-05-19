package unit

import (
	"net/url"
	"strconv"
)

// OcpcConversionInfosRequest 获取可选的深度转化目标 API Request
type OcpcConversionInfosRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignType 计划类型
	CampaignType int `json:"campaign_type,omitempty"`
	// AppID 应用id campaign_type=2/7该参数必填
	AppID uint64 `json:"app_id,omitempty"`
}

// Url implement GetRequest interface
func (r OcpcConversionInfosRequest) Url() string {
	return "v1/ad_unit/ocpc/conversion_infos"
}

// Encode implement GetRequest interface
func (r OcpcConversionInfosRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("campaign_type", strconv.Itoa(r.CampaignType))
	if r.AppID > 0 {
		values.Set("app_id", strconv.FormatUint(r.AppID, 10))
	}
	return values.Encode()
}
