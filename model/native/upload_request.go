package native

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// UploadRequest 本地视频上传接口 API Request
type UploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Photo 视频file
	Photo *model.UploadField `json:"photo,omitempty"`
	// BlobStoreKey 填写该字段时需要先使用文件流式上传或分片上传。
	BlobStoreKey string `json:"blob_store_key,omitempty"`
	// Signature 视频 md5 值，填写blob_store_key时需要填写。
	Signature string `json:"signature,omitempty"`
	// ShieldBackwardSwitch 上传视频后是否自动同步至快手个人主页，false表示屏蔽，视频不可在个人主页可见，true表示不屏蔽
	ShieldBackwardSwitch bool `json:"shieldBackwardSwitch,omitempty"`
	// AuthorID 原生上传至达人的快手号
	AuthorID uint64 `json:"authorId,omitempty"`
	// PhotoName 视频名称，默认是视频 id
	// 不超过 50 字符，若不传默认为文件名称
	PhotoName string `json:"photo_name,omitempty"`
	// PhotoTag 视频标签
	// 单个标签不超过 10 字符，支持一个标签
	PhotoTag string `json:"photo_tag,omitempty"`
	// PhotoCaption 视频描述
	PhotoCaption string `json:"photoCaption,omitempty"`
	// NativePlcSwitch 挂载plc组件（建议设置为true，有机会获得更多流量曝光）
	NativePlcSwitch bool `json:"nativePlcSwitch,omitempty"`
	// Sync 用同步/异步方式上传视频
	// 1代表同步 0代表异步，默认同步上传。 使用异步上传存在两个风险： 1、异步上传可能出现视频转码失败case。 2、如果创建物料时，视频解码失败或者解码中，生成物料会失败。 建议同步上传。
	Sync int `json:"sync,omitempty"`
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
	ret := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
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
			Key:   "nativePlcSwitch",
			Value: nativePlcSwitch,
		},
		{
			Key:   "sync",
			Value: strconv.Itoa(r.Sync),
		},
	}
	if r.Photo != nil && r.Photo.Reader != nil {
		ret = append(ret,
			model.UploadField{
				Key:    "photo",
				Value:  fileName,
				Reader: r.Photo.Reader,
			},
		)
	} else if r.BlobStoreKey != "" {
		ret = append(ret,
			model.UploadField{
				Key:   "blob_store_key",
				Value: r.BlobStoreKey,
			},
			model.UploadField{
				Key:   "signature",
				Value: r.Signature,
			},
		)
	}
	if r.PhotoName != "" {
		ret = append(ret,
			model.UploadField{
				Key:   "photo_name",
				Value: r.PhotoName,
			},
		)
	}
	if r.PhotoTag != "" {
		ret = append(ret,
			model.UploadField{
				Key:   "photo_tag",
				Value: r.PhotoTag,
			},
		)
	}
	if r.PhotoCaption != "" {
		ret = append(ret,
			model.UploadField{
				Key:   "photo_caption",
				Value: r.PhotoCaption,
			},
		)
	}
	return ret
}
