package app

import "github.com/bububa/kwai-marketing-api/model"

// RetryBuildSubPackageRequest 分包失败，重新构建 API Request
type RetryBuildSubPackageRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用ID
	AppID []uint64 `json:"app_id,omitempty"`
}

// Url implement PostRequest interface
func (r RetryBuildSubPackageRequest) Url() string {
	return "gw/dsp/appcenter/app/retryBuildSubPackage"
}

// Encode implement PostRequest interface
func (r RetryBuildSubPackageRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// RetryBuildSubPackageResponse 分包失败，重新构建 API Response
type RetryBuildSubPackageResponse struct {
	// RetryCnt 本次发起重新构建，应用分包的数量。
	// cnt=0表示没有需要重建的分包，cnt=x表示对x个分包发起了重建
	RetryCnt int `json:"retry_cnt,omitempty"`
}
