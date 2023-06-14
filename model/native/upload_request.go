package native

import (
	"github.com/bububa/kwai-marketing-api/model"
	"strconv"
)

type UploadRequest struct {
	Photo                *model.UploadField `json:"photo"`                  // 视频file
	ShieldBackwardSwitch bool               `json:"shield_backward_switch"` // 上传视频后是否自动同步至快手个人主页，false表示屏蔽，视频不可在个人主页可见，true表示不屏蔽
	AuthorID             uint64             `json:"author_id"`              // 原生上传至达人的快手号
}

// Url implement UploadRequest interface
func (r UploadRequest) Url() string {
	return "gw/dsp/v1/photo/upload"
}

// Encode implenent UploadRequest interface
func (r UploadRequest) Encode() []model.UploadField {
	fileName := r.Photo.Value
	if fileName == "" {
		fileName = "file"
	}
	shieldBackwardSwitch := "false"
	if r.ShieldBackwardSwitch {
		shieldBackwardSwitch = "true"
	}
	return []model.UploadField{
		{
			Key:   "author_id",
			Value: strconv.FormatUint(r.AuthorID, 10),
		},
		{
			Key:   "shield_backward_switch",
			Value: shieldBackwardSwitch,
		},
		{
			Key:    "file",
			Value:  fileName,
			Reader: r.Photo.Reader,
		},
	}
}
