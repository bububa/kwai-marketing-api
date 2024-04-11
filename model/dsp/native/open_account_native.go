package native

import "github.com/bububa/kwai-marketing-api/model"

// OpenAccountNativeRequest 开启原生扩量开关 API Request
type OpenAccountNativeRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OpenAccountNative 开启/关闭原生扩量开关
	// 1开启 0关闭
	OpenAccountNative int `json:"open_account_native"`
}

// Url implement PostRequest interface
func (r OpenAccountNativeRequest) Url() string {
	return "gw/dsp/native/openAccountNative"
}

// Encode implement PostRequest interface
func (r OpenAccountNativeRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// OpenAccountNativeResponse 开启原生扩量开关 API Response
type OpenAccountNativeResponse struct {
	// OpenAccountNative 操作结果 true:成功 false：失败
	OpenAccountNative bool `json:"open_account_native,omitempty"`
}
