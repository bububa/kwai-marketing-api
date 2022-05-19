package dmp

import (
	"encoding/json"
)

// PopulationUploadRequest 人群包上传接口
type PopulationUploadRequestV2 struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 人群数据类型; 1：IMEI；2：IDFA；3：IMEI_MD5；4：IDFA_MD5；5：手机号-MD5；7：OAID；8：OAID_MD5
	Type int `json:"type,omitempty"`
	// OrientationName 人群包名称; 不能大于20个字符，人群包名称不得重复
	OrientationName string `json:"orientation_name,omitempty"`
	// File 请求的文件; 最大支持500MB
	FilePaths []string `json:"file_paths,omitempty"`
}

// Url implement UploadRequest interface
func (r PopulationUploadRequestV2) Url() string {
	return "gw/dmp/v2/dmp/population/upload"
}
func (r PopulationUploadRequestV2) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
