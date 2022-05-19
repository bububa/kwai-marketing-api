package file

import "encoding/json"

// AdVideoTagDeleteRequest 视频库-删除视频标签 API Request
type AdVideoTagDeleteRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 ids，不超过 10 个
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// PhotoTag 视频标签 单个标签长度不能超过 10，只支持一个标签
	PhotoTag []string `json:"photo_tag,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoTagDeleteRequest) Url() string {
	return "v1/file/ad/video/tag/delete"
}

// Encode implement PostRequest interface
func (r AdVideoTagDeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
