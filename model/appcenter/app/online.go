package app

import "github.com/bububa/kwai-marketing-api/model"

// OnlineRequest 【应用中心】应用上架 API Request
type OnlineRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageIDs 应用包ID
	PackageIDs []uint64 `json:"package_ids,omitempty"`
}

// Url implement PostRequest interface
func (r OnlineRequest) Url() string {
	return "gw/dsp/appcenter/app/online"
}

// Encode implement PostRequest interface
func (r OnlineRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// OnlineOfflineResponse 【应用中心】应用上/下架 API Response
type OnlineOfflineResponse struct {
	// Result 上架结果
	Result bool `json:"result,omitempty"`
}
