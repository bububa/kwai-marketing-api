package native

import (
	"github.com/bububa/kwai-marketing-api/model"
	"strconv"
)

type UploadRequest struct {
	AdvertiserId         uint64             `json:"advertiser_id"`
	Photo                *model.UploadField `json:"photo"`                // 视频file
	ShieldBackwardSwitch bool               `json:"shieldBackwardSwitch"` // 上传视频后是否自动同步至快手个人主页，false表示屏蔽，视频不可在个人主页可见，true表示不屏蔽
	AuthorID             uint64             `json:"authorId"`             // 原生上传至达人的快手号
	NativePlcSwitch      bool               `json:"nativePlcSwitch"`
	PhotoCaption         string             `json:"photoCaption"`
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
	nativePlcSwitch := "false"
	if r.NativePlcSwitch {
		nativePlcSwitch = "true"
	}
	return []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserId, 10),
		},
		{
			Key:   "authorId",
			Value: strconv.FormatUint(r.AuthorID, 10),
		},
		{
			Key:   "shieldBackwardSwitch",
			Value: shieldBackwardSwitch,
		},
		{
			Key:    "photo",
			Value:  fileName,
			Reader: r.Photo.Reader,
		},
		{
			Key:   "nativePlcSwitch",
			Value: nativePlcSwitch,
		},
		{
			Key:   "photoCaption",
			Value: r.PhotoCaption,
		},
	}
}
