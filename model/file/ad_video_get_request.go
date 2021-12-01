package file

import "encoding/json"

// AdVideoGetRequest 获取视频信息get接口 API Request
type AdVideoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频ID集
	PhotoIDs []string `json:"photo_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoGetRequest) Url() string {
	return "v1/file/ad/video/get"
}

// Encode implement PostRequest interface
func (r AdVideoGetRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
