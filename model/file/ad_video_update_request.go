package file

import "encoding/json"

// AdVideoUpdateRequest 视频库-批量更新视频功能 API Request
type AdVideoUpdateRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 ids，不超过 100 个
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// PhotoName 视频名称; photo_name 和 photo_tags 填其一即可。photo_name长度不得超过100字符
	PhotoName string `json:"photo_name,omitempty"`
	// PhotoTag 视频标签 单个标签长度不能超过 10，只支持一个标签
	PhotoTag []string `json:"photo_tag,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoUpdateRequest) Url() string {
	return "v1/file/ad/video/update"
}

// Encode implement PostRequest interface
func (r AdVideoUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
