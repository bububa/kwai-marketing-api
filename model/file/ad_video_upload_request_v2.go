package file

import (
	"strconv"

	"git.gametaptap.com/tapad/github/kwai-marketing-api/model"
)

// AdVideoUploadRequestV2 上传视频v2接口 API Request
type AdVideoUploadRequestV2 struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// File 视频文件; 只支持mp4格式，详细要求见附录
	File *model.UploadField `json:"file,omitempty"`
	// Type 上传视频类型; 0：发现页信息流视频，默认值 1：信息流竖版视频 2：信息流横版视频 3：后贴片竖版视频（已下线） 4：后贴片横版视频（已下线）
	Type int `json:"type"`
	// Signature 视频md5值
	Signature string `json:"signature,omitempty"`
	// PhotoName 视频名称，默认是视频 id; 不超过 50 字符，若不传默认为文件名称
	PhotoName string `json:"photo_name,omitempty"`
	// PhotoTag 视频标签; 单个标签不超过 10 字符，支持一个标签
	PhotoTag []string `json:"photo_tag,omitempty"`
	// Sync 用同步/异步方式上传视频; 0:(默认)以异步方式上传，不需要同步等待，上传较快。1:(选填)同步方式上传，上传较慢
	Sync int `json:"sync,omitempty"`
	// ShieldBackwardSwitch 上传视频后是否自动同步至快手个人主页; false表示后端屏蔽，视频不可profile页可见， true表示不屏蔽。
	ShieldBackwardSwitch bool `json:"shield_backward_switch,omitempty"`
}

// Url implement UploadRequest interface
func (r AdVideoUploadRequestV2) Url() string {
	return "v2/file/ad/video/upload"
}

// Encode implenent UploadRequest interface
func (r AdVideoUploadRequestV2) Encode() []model.UploadField {
	fileName := r.File.Value
	if fileName == "" {
		fileName = "file"
	}
	shieldBackwardSwitch := "false"
	if r.ShieldBackwardSwitch {
		shieldBackwardSwitch = "true"
	}
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatInt(r.AdvertiserID, 10),
		},
		{
			Key:   "type",
			Value: strconv.Itoa(r.Type),
		},
		{
			Key:   "sync",
			Value: strconv.Itoa(r.Sync),
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
	if r.PhotoName != "" {
		fields = append(fields, model.UploadField{
			Key:   "photo_name",
			Value: r.PhotoName,
		})
	}
	if len(r.PhotoTag) > 0 {
		for _, tag := range r.PhotoTag {
			fields = append(fields, model.UploadField{
				Key:   "[]photo_tag",
				Value: tag,
			})
		}
	}
	return fields
}
