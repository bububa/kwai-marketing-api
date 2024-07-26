package app

import "github.com/bububa/kwai-marketing-api/model"

// CreateIosRequest 创建 iOS 应用 API Request
type CreateIosRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppIconURL 应用图标，可不填，如果填写则覆盖应用解析出的应用图标，可通过通过上传图片 API 获取链接。
	AppIconURL string `json:"app_icon_url,omitempty"`
	// IosDownloadURL App Store 应用下载链接
	IosDownloadURL string `json:"ios_download_url,omitempty"`
}

// Url implements PostRequest interface
func (r CreateIosRequest) Url() string {
	return "gw/dsp/appcenter/app/create/ios"
}

// Encode implements PostRequest interface
func (r CreateIosRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
