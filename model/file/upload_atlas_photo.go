package file

import "github.com/bububa/kwai-marketing-api/model"

// UploadAtlasPhotoRequest 上传图文视频 API Request
type UploadAtlasPhotoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// 图片ID
	PicIDs []string `json:"pic_ids,omitempty"`
	// AudioBsKey 音频bs_key
	AudioBsKey string `json:"audio_bs_key,omitempty"`
	// ShieldBackwordSwitch 上传视频后是否自动同步至快手个人主页, 默认false
	ShieldBackwordSwitch bool `json:"shield_backword_switch,omitempty"`
	// WaitForTranscode 用同步/异步方式上传视频,默认false
	WaitForTranscode bool `json:"wait_for_transcode,omitempty"`
}

// Url implement PostRequest interface
func (r UploadAtlasPhotoRequest) Url() string {
	return "gw/dsp/upload/atlasPhoto"
}

// Encode implement PostRequest interface
func (r UploadAtlasPhotoRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UploadAtlasPhotoResponse 上传图文视频 API Response
type UploadAtlasPhotoResponse struct {
	// PhotoID 图文视频ID
	PhotoID string `json:"photo_id,omitempty"`
}
