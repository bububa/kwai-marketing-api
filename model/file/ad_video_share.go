package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoShareRequest 视频库-推送视频 API Request
type AdVideoShareRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 ids，不超过 10 个
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// ShareAdvertiserIDs 推送账户; 同一资质下其他 user_id 下的广告主 ids，不超过10个，同一个快手 user_id 下视频是共享的（无需推送）， 视频共享只能是同一个资质主体下不同的 user_id 的广告主可以共享
	ShareAdvertiserIDs []uint64 `json:"share_advertiser_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoShareRequest) Url() string {
	return "v1/file/ad/video/share"
}

// Encode implement PostRequest interface
func (r AdVideoShareRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdVideoShareResponse 视频库-推送视频 API Response
type AdVideoShareResponse struct {
	// Details Response details
	Details []AdVideoShareDetail `json:"details,omitempty"`
}

// AdVideoShareDetail 视频库-推送视频 API Response Detail
type AdVideoShareDetail struct {
	// AdvertiserID 账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoID 分享生成新的photoId
	PhotoID uint64 `json:"photo_id,omitempty"`
	// OriginalPhotoID 原始photoId
	OriginalPhotoID uint64 `json:"original_photo_id,omitempty"`
}
