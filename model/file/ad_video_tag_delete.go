package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoTagDeleteRequest 视频库-删除视频标签 API Request
type AdVideoTagDeleteRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoTag 视频标签 单个标签长度不能超过 10，只支持一个标签
	PhotoTag []string `json:"photo_tag,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoTagDeleteRequest) Url() string {
	return "v1/file/ad/video/tag/delete"
}

// Encode implement PostRequest interface
func (r AdVideoTagDeleteRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
