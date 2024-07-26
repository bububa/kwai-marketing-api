package app

import "github.com/bububa/kwai-marketing-api/model"

// DetailRequest 【应用中心】获取应用详情 API Request
type DetailRequest struct {
	// AdvertiserID	Long		必填	广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用包id
	// 仅支持应用母包
	PackageID uint64 `json:"package_id,omitempty"`
}

// Url implements PostRequest interface
func (r DetailRequest) Url() string {
	return "gw/dsp/appcenter/app/detail"
}

// Encode implement PostRequest interface
func (r DetailRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
