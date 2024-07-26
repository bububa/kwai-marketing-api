package upload

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// PicRequest 上传图片 API Request
type PicRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File 图片文件
	File *model.UploadField `json:"file,omitempty"`
	// Type 图片类型：1=应用图标；2=应用图片, 5=单机承诺函
	// 应用图标：尺寸：450 x 450，小于1MB；格式为 PNG/JPG/JPEG。应用图片：宽高比9:20，宽：大于等于720，高：大于等于1280，小于2 MB；格式为 PNG/JPG/JPEG。单机承诺函：小于10MB；格式为 PNG/JPG/JPEG。
	Type int `json:"type,omitempty"`
}

// Url implement UploadRequest interface
func (r PicRequest) Url() string {
	return "gw/dsp/appcenter/upload/pic"
}

// Encode implenent UploadRequest interface
func (r PicRequest) Encode() []model.UploadField {
	fileName := r.File.Value
	if fileName == "" {
		fileName = "file"
	}
	return []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
		{
			Key:   "type",
			Value: strconv.Itoa(r.Type),
		},
		{
			Key:    "file",
			Value:  fileName,
			Reader: r.File.Reader,
		},
	}
}

// PicResponse 上传图片 API Response
type PicResponse struct {
	// URL 所上传图片的快手CDN链接
	URL string `json:"url,omitempty"`
}
