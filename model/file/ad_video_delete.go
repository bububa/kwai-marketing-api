package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoDeleteRequest 批量删除视频 API Request
type AdVideoDeleteRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 ids，不超过 100 个
	PhotoIDs []string `json:"photo_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoDeleteRequest) Url() string {
	return "gw/dsp/file/ad/video/delete"
}

// Encode implement PostRequest interface
func (r AdVideoDeleteRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdVideoDeleteResponse 批量删除视频 API Response
type AdVideoDeleteResponse struct {
	// PhotoIDs 视频 ids
	PhotoIDs []string `json:"photo_ids,omitempty"`
}
