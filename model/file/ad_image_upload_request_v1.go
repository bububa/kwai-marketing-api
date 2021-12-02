package file

import (
	"strconv"

	"git.gametaptap.com/tapad/github/kwai-marketing-api/model"
)

// AdImageUploadRequestV1 上传图片v1接口 API Request
type AdImageUploadRequestV1 struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// File 图片文件
	File *model.UploadField `json:"file,omitempty"`
	// Type 上传图片类型; 0：上传封面图片（仅1.0版本使用），1：上传app_icon图片，2：广告组设置中scene_id为7时的封面图片；3：广告组设置中scene_id为3时的素材图片（后贴片样式已下线）
	Type int `json:"type"`
}

// Url implement UploadRequest interface
func (r AdImageUploadRequestV1) Url() string {
	return "v1/file/ad/image/upload"
}

// Encode implenent UploadRequest interface
func (r AdImageUploadRequestV1) Encode() []model.UploadField {
	fileName := r.File.Value
	if fileName == "" {
		fileName = "file"
	}
	return []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatInt(r.AdvertiserID, 10),
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
