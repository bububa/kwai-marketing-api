package app

import "github.com/bububa/kwai-marketing-api/model"

// UpdateAndroidRequest 更新 Android 应用 API Request
type UpdateAndroidRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppInfo 应用数据
	AppInfo *App `json:"app_info,omitempty"`
	// AppPrivacyInfo 隐私声明数据
	AppPrivacyInfo *AppPrivacyInfo `json:"app_privacy_info,omitempty"`
	// PackageInfo 应用包数据
	PackageInfo *PackageInfo `json:"package_info,omitempty"`
}

// Url implements PostRequest interface
func (r UpdateAndroidRequest) Url() string {
	return "gw/dsp/appcenter/app/update/android"
}

// Encode implements PostRequest interface
func (r UpdateAndroidRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
