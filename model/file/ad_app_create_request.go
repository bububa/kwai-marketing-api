package file

import (
	"strconv"

	"github.com/bububa/kwai-marketing-api/model"
)

// AdAppCreateRequest 创建应用 API Request
type AdAppCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File 应用安装包（apk）; 包体限制：1G 以内；当platform为1时，可自行选择上传apk，如果同时传了url，以file字段值为主
	File *model.UploadField
	// AppVersion 应用标记; 不能超过100字符，同一账户下应用标记不能重复；样例：快影-3.0.0.0103
	AppVersion string `json:"app_version,omitempty"`
	// AppName 应用名称; 不能超过100字符
	AppName string `json:"app_name,omitempty"`
	// ImageToken 应用图标的image_token; 图片token，可通过上传图片接口获取；platform为1/3时必填，请上传png/jpg/jpeg图片，尺寸450*450，小于100KB
	ImageToken string `json:"image_token,omitempty"`
	// PackageName 应用包名; platform为1（Android应用下载）必填，其它类型不用填写，不能超过 100 字符
	PackageName string `json:"package_name,omitempty"`
	// Platform 应用类型; 1：Android应用下载，2：Android网页游戏，3：iOS应用下载，4：iOS网页游戏
	Platform int `json:"platform,omitempty"`
	// URL 应用地址
	URL string `json:"url,omitempty"`
	// UseSDK 是否接入快手广告监测SDK; 0：未接入，1：已接入
	UseSDK int `json:"use_sdk"`
	// AppPrivacyUrl app隐私政策链接，需与app相关，该字段会经过审核; 安卓类应用必填
	AppPrivacyUrl string `json:"app_privacy_url,omitempty"`
	// PermissionInformation 权限信息，请通过应用权限信息列表接口获取信息
	PermissionInformation string `json:"permission_information,omitempty"`
	// RealAppVersion 真实版本号
	RealAppVersion string `json:"real_app_version,omitempty"`
	// AppIconUrl 应用包大小
	PackageSize int64 `json:"package_size,omitempty"`
	// AppDetailImageToken app应用详情图片
	AppDetailImageToken string `json:"app_detail_image_token,omitempty"`
}

// Url implement UploadRequest interface
func (r AdAppCreateRequest) Url() string {
	return "v1/file/ad/app/create"
}

// Encode implement UploadRequest interface
func (r AdAppCreateRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
		{
			Key:   "app_version",
			Value: r.AppVersion,
		},
		{
			Key:   "app_name",
			Value: r.AppName,
		},
		{
			Key:   "platform",
			Value: strconv.Itoa(r.Platform),
		},
		{
			Key:   "use_sdk",
			Value: strconv.Itoa(r.UseSDK),
		},
	}
	if r.File != nil {
		filename := r.File.Value
		if filename == "" {
			filename = "file"
		}
		fields = append(fields, model.UploadField{
			Key:    "file",
			Value:  filename,
			Reader: r.File.Reader,
		})
	}
	if r.URL != "" {
		fields = append(fields, model.UploadField{
			Key:   "url",
			Value: r.URL,
		})
	}
	if r.AppDetailImageToken != "" {
		fields = append(fields, model.UploadField{
			Key:   "app_detail_image_token",
			Value: r.AppDetailImageToken,
		})
	}
	if r.ImageToken != "" {
		fields = append(fields, model.UploadField{
			Key:   "image_token",
			Value: r.ImageToken,
		})
	}
	if r.AppPrivacyUrl != "" {
		fields = append(fields, model.UploadField{
			Key:   "app_privacy_url",
			Value: r.AppPrivacyUrl,
		})
	}
	if r.PermissionInformation != "" {
		fields = append(fields, model.UploadField{
			Key:   "permission_information",
			Value: r.PermissionInformation,
		})
	}
	if r.RealAppVersion != "" {
		fields = append(fields, model.UploadField{
			Key:   "real_app_version",
			Value: r.RealAppVersion,
		})
	}
	if r.PackageSize != 0 {
		fields = append(fields, model.UploadField{
			Key:   "package_size",
			Value: strconv.FormatInt(r.PackageSize, 10),
		})
	}
	if r.PackageName != "" {
		fields = append(fields, model.UploadField{
			Key:   "package_name",
			Value: r.PackageName,
		})
	}

	return fields
}
