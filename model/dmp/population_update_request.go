package dmp

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// PopulationUpdateRequest 人群包更新接口
type PopulationUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Type 人群数据类型; 1：IMEI；2：IDFA；3：IMEI_MD5；4：IDFA_MD5；5：手机号-MD5；7：OAID；8：OAID_MD5
	Type int `json:"type,omitempty"`
	// OrientationID 人群包id
	OrientationID int64 `json:"orientation_id,omitempty"`
	// OperationType 修改操作类型; 1：APPEND（增量更新）3：DELETE（缩量更新）
	OperationType int `json:"operation_type,omitempty"`
	// File 请求的文件; 最大支持500MB
	File *model.UploadField `json:"file,omitempty"`
}

// Url implement UploadRequest interface
func (r PopulationUpdateRequest) Url() string {
	return "v1/dmp/population/update"
}

// Encode implement UploadRequest interface
func (r PopulationUpdateRequest) Encode() []model.UploadField {
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
			Key:   "orientation_id",
			Value: strconv.FormatInt(r.OrientationID, 10),
		},
		{
			Key:   "operation_type",
			Value: strconv.Itoa(r.OperationType),
		},
		{
			Key:    "file",
			Value:  fileName,
			Reader: r.File.Reader,
		},
	}
}
