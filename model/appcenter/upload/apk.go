package upload

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// ApkRequest 上传 APK 文件 API Request
type ApkRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File APK 文件
	File *model.UploadField `json:"file,omitempty"`
}

// Url implement UploadRequest interface
func (r ApkRequest) Url() string {
	return "gw/dsp/appcenter/upload/apk"
}

// Encode implenent UploadRequest interface
func (r ApkRequest) Encode() []model.UploadField {
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
			Key:    "file",
			Value:  fileName,
			Reader: r.File.Reader,
		},
	}
}

// ApkResponse 上传 APK 文件 API Response
type ApkResponse struct {
	// BlobStoreKey APK 文件在快手的存储 Key
	BlobStoreKey string `json:"blob_store_key,omitempty"`
}
