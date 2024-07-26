package app

import "github.com/bububa/kwai-marketing-api/model"

// IosUpdateRequest 【应用中心】iOS 应用上报更新
type IosUpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// IosAppID 苹果商店 iOS App Id
	IosAppID uint64 `json:"ios_app_id,omitempty"`
	// PackageID 应用包ID
	PackageID uint64 `json:"package_id,omitempty"`
}

// Url implement PostRequest interface
func (r IosUpdateRequest) Url() string {
	return "gw/dsp/appcenter/app/ios/update"
}

// Encode implement PostRequest interface
func (r IosUpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// IosUpdateResponse【应用中心】iOS 应用上报更新 API Response
type IosUpdateResponse struct {
	// Result 上报更新结果
	Result bool `json:"result,omitempty"`
}
