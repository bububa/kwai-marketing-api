package file

import "encoding/json"

// AdVideoRelateCreativesRequest 视频关联创意数查询
type AdVideoRelateCreativesRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频id; 最大20个，可以动态配置
	PhotoIds []string `json:"photo_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoRelateCreativesRequest) Url() string {
	return "v1/file/ad/video/relate/creatives"
}

// Encode implement PostRequest interface3
func (r AdVideoRelateCreativesRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
