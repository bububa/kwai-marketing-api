package app

import "github.com/bububa/kwai-marketing-api/model"

// OfflineAppStoresRequest 【应用中心】应用商店上下架 API Request
type OfflineAppStoresRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppIDs 应用ID
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// OfflineStores 应用商店
	OfflineStores []string `json:"offline_stores,omitempty"`
}

// Url implement PostRequest interface
func (r OfflineAppStoresRequest) Url() string {
	return "gw/dsp/appcenter/app/offline/appstores"
}

// Encode implement PostRequest interface
func (r OfflineAppStoresRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
