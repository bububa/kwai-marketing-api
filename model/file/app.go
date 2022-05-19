package file

// App 应用
type App struct {
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// URL 应用地址
	URL string `json:"url,omitempty"`
	// AppVersion 应用标记
	AppVersion string `json:"app_version,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// AppIconUrl 应用图标链接
	AppIconUrl string `json:"app_icon_url,omitempty"`
	// ImageToken 应用图标的image_token
	ImageToken string `json:"image_token,omitempty"`
	// PacakgeName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// Platform 应用类型
	Platform int `json:"platform,omitempty"`
	// UpdateTime 更新时间; 单位：毫秒
	UpdateTime int64 `json:"update_time,omitempty"`
	// UseSDK 是否接入快手广告监测SDK; 0：未接入，1：已接入
	UseSDK int `json:"use_sdk,omitempty"`
	// AppPrivacyUrl app隐私政策链接，需与app相关，该字段会经过审核; 安卓类应用必填
	AppPrivacyUrl string `json:"app_privacy_url,omitempty"`
	// ScanStatus 应用安全扫描状态; 1-扫描中，2-成功，3-失败，4-失败重试中
	ScanStatus int `json:"scan_status,omitempty"`
	// PermissionInformation 权限信息，请通过应用权限信息列表接口获取信息
	PermissionInformation []int `json:"permission_information,omitempty"`
	// RealAppVersion 真实版本号
	RealAppVersion string `json:"real_app_version,omitempty"`
	// PackageSize 应用包大小
	PackageSize int64 `json:"package_size,omitempty"`
	// AppDetailImageToken app应用详情图片
	AppDetailImageToken string `json:"app_detail_image_token,omitempty"`
}
