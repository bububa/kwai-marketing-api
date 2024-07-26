package subpkg

// SubPackage 应用分包
type SubPackage struct {
	// AccountID	Long		账号ID
	AccountII uint64 `json:"account_id,omitempty"`
	// AppDetailImg	String		应用详情图片
	AppDetailImg string `json:"app_detail_img,omitempty"`
	// AppIconURL	String		应用图标链接
	AppIconURL string `json:"app_icon_url,omitempty"`
	// AppID	Long		应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// AppPrivacyURL	String		应用隐私政策链接
	AppPrivacyURL string `json:"app_privacy_url,omitempty"`
	// ChannelID String 渠道号(分包号)
	ChannelID string `json:"channel_id,omitempty"`
	// IosAppID	String		解析出的iosAppID
	IosAppID string `json:"ios_app_id,omitempty"`
	// OfflineAppStores	String		下架的应用商店	"huawei","oppo","vivo","xiaomi","meizu","smartisan"
	OfflineAppStores string `json:"offline_app_stores,omitempty"`
	// PackageID	Long		应用包ID
	PackageID uint64 `json:"package_id,omitempty"`
	// PackageName	String		应用包名
	PackageName string `json:"package_name,omitempty"`
	// PackageSize	Long		应用包大小
	PackageSize int64 `json:"package_size,omitempty"`
	// ParentPackageID 分包的母包ID
	ParentPackageID uint64 `json:"parent_package_id,omitempty"`
	// PermissionInformation	int[]		权限信息ID列表
	PermissionInformation []int `json:"permission_information,omitempty"`
	// Platform	String		android或ios
	Platform string `json:"platform,omitempty"`
	// RealAppName	String		应用名称
	RealAppName string `json:"real_app_name,omitempty"`
	// RealAppVersion	String		应用版本信息
	RealAppVersion string `json:"real_app_version,omitempty"`
	// SourceType	Integer		应用来源	1-我创建的 2-共享给我的
	SourceType int `json:"source_type,omitempty"`
	// SubPackageStatus 应用分包状态
	// 1-审核中，2-审核失败，3-待发布，4-已发布，5-已下架 6-创建中，7-更新中，8-构建失败
	SubPackageStatus int `json:"sub_package_status,omitempty"`
	// UpdateTime	Long		更新时间	单位：毫秒
	UpdateTime int64 `json:"update_time,omitempty"`
	// DeleteTime 删除时间
	// 仅分包回收站列表时有效，表示应用分包的删除时间
	DeleteTime int64 `json:"delete_time,omitempty"`
	// URL	String		应用下载地址
	URL string `json:"url,omitempty"`
	// UseSDK	Integer		是否接入快手广告监测SDK	0-未接入，1-已接入
	UseSDK int `json:"use_sdk,omitempty"`
	// VersionCode
	VersionCode int `json:"version_code,omitempty"`
	// BuildStatus 构建状态 0-创建中，1-构建中，2-构建成功，3-构建失败
	BuildStatus int `json:"build_status,omitempty"`
	// Description 分包备注
	Description string `json:"description,omitempty"`
	// CanUpdate 是否可更新
	// 仅分包管理列表时有效，表示应用分包是否可以更新
	CanUpdate bool `json:"can_update,omitempty"`
	// CanRecycle 是否可恢复
	// 仅分包回收站列表时有效，表示应用分包是否可以恢复
	CanRecycle bool `json:"can_recycle,omitempty"`
}
