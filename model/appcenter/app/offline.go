package app

import "github.com/bububa/kwai-marketing-api/model"

// OfflineRequest 【应用中心】应用上架 API Request
type OfflineRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageIDs 应用包ID
	PackageIDs []uint64 `json:"package_ids,omitempty"`
}

// Url implement PostRequest interface
func (r OfflineRequest) Url() string {
	return "gw/dsp/appcenter/app/offline"
}

// Encode implement PostRequest interface
func (r OfflineRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
