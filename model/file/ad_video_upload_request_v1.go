package file

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoUploadRequestV1 上传视频v1接口 API Request
type AdVideoUploadRequestV1 struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// File 视频文件; 只支持mp4格式，详细要求见附录
	File *model.UploadField `json:"file,omitempty"`
	// Type 上传视频类型; 0：发现页信息流视频，默认值 1：信息流竖版视频 2：信息流横版视频 3：后贴片竖版视频（已下线） 4：后贴片横版视频（已下线）
	Type int `json:"type"`
	// ShieldBackwardSwitch 上传视频后是否自动同步至快手个人主页; false表示后端屏蔽，视频不可profile页可见， true表示不屏蔽。
	ShieldBackwardSwitch bool `json:"shield_backward_switch,omitempty"`
}

// Url implement UploadRequest interface
func (r AdVideoUploadRequestV1) Url() string {
	return "v1/file/ad/video/upload"
}

// Encode implenent UploadRequest interface
func (r AdVideoUploadRequestV1) Encode() []model.UploadField {
	fileName := r.File.Value
	if fileName == "" {
		fileName = "file"
	}
	shieldBackwardSwitch := "false"
	if r.ShieldBackwardSwitch {
		shieldBackwardSwitch = "true"
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
			Key:   "shield_backward_switch",
			Value: shieldBackwardSwitch,
		},
		{
			Key:    "file",
			Value:  fileName,
			Reader: r.File.Reader,
		},
	}
}
