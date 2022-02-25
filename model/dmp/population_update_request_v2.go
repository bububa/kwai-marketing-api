package dmp

import (
	"encoding/json"
)

// PopulationUpdateRequestv2 人群包更新接口
type PopulationUpdateRequestv2 struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Type 人群数据类型; 1：IMEI；2：IDFA；3：IMEI_MD5；4：IDFA_MD5；5：手机号-MD5；7：OAID；8：OAID_MD5
	Type int `json:"type,omitempty"`
	// OrientationID 人群包id
	OrientationID int64 `json:"orientation_id,omitempty"`
	// OperationType 修改操作类型; 1：APPEND（增量更新）3：DELETE（缩量更新）
	OperationType int `json:"operation_type,omitempty"`
	// File 请求的文件; 通过调用【文件上传接口】得到的文件路径，每次最多填写10个、总量<3G的文件
	FilePaths []string `json:"file_paths,omitempty"`
}

// Url implement UploadRequest interface
func (r PopulationUpdateRequestv2) Url() string {
	return "gw/dmp/v2/dmp/population/update"
}

// Encode implement UploadRequest interface
func (r PopulationUpdateRequestv2) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
