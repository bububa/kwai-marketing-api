package subpkg

import "github.com/bububa/kwai-marketing-api/model"

// DescriptionRequest 【应用中心】修改应用分包备注 API Request
type DescriptionRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用分包id
	PackageID uint64 `json:"package_id,omitempty"`
	// Description 分包的备注信息	当应用分包处于构建中或更新中时，不可对分包修改备注
	Description string `json:"description,omitempty"`
}

// Url implement PostRequest interface
func (r DescriptionRequest) Url() string {
	return "gw/dsp/appcenter/subpkg/description"
}

// Encode implement PostRequest interface
func (r DescriptionRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
