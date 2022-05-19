package dmp

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// PopulationUploadRequest 人群包上传接口
type PopulationUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 人群数据类型; 1：IMEI；2：IDFA；3：IMEI_MD5；4：IDFA_MD5；5：手机号-MD5；7：OAID；8：OAID_MD5
	Type int `json:"type,omitempty"`
	// OrientationName 人群包名称; 不能大于20个字符，人群包名称不得重复
	OrientationName string `json:"orientation_name,omitempty"`
	// File 请求的文件; 最大支持500MB
	File *model.UploadField `json:"file,omitempty"`
}

// Url implement UploadRequest interface
func (r PopulationUploadRequest) Url() string {
	return "v1/dmp/population/upload"
}

// Encode implement UploadRequest interface
func (r PopulationUploadRequest) Encode() []model.UploadField {
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
			Key:   "orientation_name",
			Value: r.OrientationName,
		},
		{
			Key:    "file",
			Value:  fileName,
			Reader: r.File.Reader,
		},
	}
}
