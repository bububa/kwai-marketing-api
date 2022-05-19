package file

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// AdImageUploadRequestV2 上传图片v2接口 API Request
type AdImageUploadRequestV2 struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 上传图片类型; 默认为2。1：上传app_icon图片； 2：广告组设置中scene_id为7时的封面图片； 4：广告组设置中scene_id为3时的便利贴广告素材图片；5：联盟图片素材； 6：横版图片；7：小图（组图与小图对于格式要求一致，只不过组图传三个）；高级创意图片：100:图片卡片 ；101:多利益卡-图文 ；102：多利益卡-多标签 ；103：电商促销样式。 要求:1.图片宽度不能小于228像素，高度不能小于150像素、2.图片宽高比为1.52:1/只支持png/jpeg/jpg格式、3.图片不能大于2M
	Type int `json:"type,omitempty"`
	// UploadType 1: 通过文件上传；2: 通过图片url上传;
	UploadType int `json:"upload_type,omitempty"`
	// File 图片文件; upload_type为1时必填，详细要求见附录
	File *model.UploadField `json:"file,omitempty"`
	// URL 图片url; upload_type为2时必填，详细要求见附录
	URL string `json:"url,omitempty"`
	// Signature 图片md5值; 用于服务端校验，当 upload_type为1时必填
	Signature string `json:"siganture,omitempty"`
}

// Url implement UploadRequest interface
func (r AdImageUploadRequestV2) Url() string {
	return "v2/file/ad/image/upload"
}

// Encode implenent UploadRequest interface
func (r AdImageUploadRequestV2) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
		{
			Key:   "type",
			Value: strconv.Itoa(r.Type),
		},
		{
			Key:   "upload_type",
			Value: strconv.Itoa(r.UploadType),
		},
	}
	switch r.UploadType {
	case 1:
		fileName := r.File.Value
		if fileName == "" {
			fileName = "file"
		}
		fields = append(fields, model.UploadField{
			Key:    "file",
			Value:  fileName,
			Reader: r.File.Reader,
		}, model.UploadField{
			Key:   "signature",
			Value: r.Signature,
		})
	case 2:
		fields = append(fields, model.UploadField{
			Key:   "url",
			Value: r.URL,
		})
	}
	return fields
}
