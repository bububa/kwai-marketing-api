package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoGetRequest 获取视频信息get接口 API Request
type AdVideoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频ID集
	PhotoIDs []string `json:"photo_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoGetRequest) Url() string {
	return "v1/file/ad/video/get"
}

// Encode implement PostRequest interface
func (r AdVideoGetRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
