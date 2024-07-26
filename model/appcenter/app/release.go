package app

import "github.com/bububa/kwai-marketing-api/model"

// ReleaseRequest 发布应用 API Request
type ReleaseRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageIDs 应用包ID【母包ID】。只有安卓审核通过的母包才可以发布，发布后的package_id才能绑定单元。
	PackageIDs []uint64 `json:"package_ids,omitempty"`
}

// Url implement PostRequest interface
func (r ReleaseRequest) Url() string {
	return "gw/dsp/appcenter/app/release"
}

// Encode implement PostRequest interface
func (r ReleaseRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
