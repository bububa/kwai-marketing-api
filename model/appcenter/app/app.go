package app

// App 应用
type App struct {
	// AccountID	Long		账号ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AppDetailImg	String		应用详情图片
	AppDetailImg string `json:"app_detail_img,omitempty"`
	// AppIconURL	String		应用图标链接
	AppIconURL string `json:"app_icon_url,omitempty"`
	// AppID	Long		应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// AppPrivacyURL	String		应用隐私政策链接
	AppPrivacyURL string `json:"app_privacy_url,omitempty"`
	// ApplyAge 使用年龄
	ApplyAge int `json:"apply_age,omitempty"`
	// AuditSerialNumber 审核序列号
	AuditSerialNumber uint64 `json:"audit_serial_number,omitempty"`
	// Category 应用类型
	// 1-软件 2-游戏
	Category int `json:"category,omitempty"`
	// Compatibility 兼容性
	Compatibility string `json:"compatibility,omitempty"`
	// ContactEmail 联系人邮箱
	ContactEmail string `json:"contact_email,omitempty"`
	// ContactName 联系人姓名
	ContactName string `json:"contact_name,omitempty"`
	// ContactTel 联系人电话
	ContactTel string `json:"contact_tel,omitempty"`
	// Description 备注
	Description string `json:"description,omitempty"`
	// Developer 开发者
	Developer string `json:"developer,omitempty"`
	// Location 开发者地区
	Location string `json:"location,omitempty"`
	// IosAppID	String		解析出的iosAppID
	IosAppID string `json:"ios_app_id,omitempty"`
	// OfflineAppStores	String		下架的应用商店	"huawei","oppo","vivo","xiaomi","meizu","smartisan"
	OfflineAppStores string `json:"offline_app_stores,omitempty"`
	// OnlineEarnType 网赚类型
	// 1-是 2-否
	OnlineEarnType int `json:"online_earn_type,omitempty"`
	// PackageID	Long		应用包ID
	PackageID uint64 `json:"package_id,omitempty"`
	// PackageName	String		应用包名
	PackageName string `json:"package_name,omitempty"`
	// PackageSize	Long		应用包大小
	PackageSize int64 `json:"package_size,omitempty"`
	// PermissionInformation	int[]		权限信息ID列表
	PermissionInformation []int `json:"permission_information,omitempty"`
	// Platform	String		android或ios
	Platform string `json:"platform,omitempty"`
	// PrivacyID 隐私ID
	PrivacyID uint64 `json:"privacy_id,omitempty"`
	// PrivacyType 隐私链接类型
	PrivacyType int `json:"privacy_type,omitempty"`
	// RealAppName	String		应用名称
	RealAppName string `json:"real_app_name,omitempty"`
	// RealAppVersion	String		应用版本信息
	RealAppVersion string `json:"real_app_version,omitempty"`
	// ReleaseType 发布类型
	// 1-手动 2-自动
	ReleaseType int `json:"release_type,omitempty"`
	// SourceType	Integer		应用来源	1-我创建的 2-共享给我的
	SourceType int `json:"source_type,omitempty"`
	// AppSource 应用创建者信息
	AppSource []AppSource `json:"app_source,omitempty"`
	// AppStatus 应用状态
	// 1-审核中 2-审核失败 3-待发布 4-已发布 5-已下架
	AppStatus int `json:"app_status,omitempty"`
	// PutStatus 投放状态
	PutStatus int `json:"put_status,omitempty"`
	// ReviewDetail 审核详情
	ReviewDetail string `json:"review_detail,omitempty"`
	// ReviewStatus 审核状态
	ReviewStatus int `json:"review_status,omitempty"`
	// SensitivePermissionDesc 敏感权限用途
	SensitivePermissionDesc string `json:"sensitive_permission_desc,omitempty"`
	// ShareAccountCount 应用共享账号个数
	ShareAccountCount int `json:"share_account_count,omitempty"`
	// TraceActivation 转化追踪状态
	TraceActivation int `json:"trace_activation,omitempty"`
	// UpdateTime	Long		更新时间	单位：毫秒
	UpdateTime int64 `json:"update_time,omitempty"`
	// URL	String		应用下载地址
	URL string `json:"url,omitempty"`
	// UseSDK	Integer		是否接入快手广告监测SDK	0-未接入，1-已接入
	UseSDK int `json:"use_sdk,omitempty"`
	// VersionCode 应用版本号
	VersionCode int `json:"version_code,omitempty"`
	// ShareType 共享类型，0-不共享，1-账号，2-主体
	ShareType int `json:"share_type,omitempty"`
	// FunctionIntroduction 安卓应用功能介绍
	FunctionIntroduction string `json:"function_introduction,omitempty"`
	// RecordNumber 备案号
	RecordNumber string `json:"record_number,omitempty"`
	// DocumentNumber 证件号码
	DocumentNumber string `json:"document_number,omitempty"`
	// ServiceCategory 服务类目
	ServiceCategory string `json:"service_category,omitempty"`
	// NetworkType 网络类型
	// 1-联网，2-单机
	NetworkType int `json:"network_type,omitempty"`
	// OfflineAppLetterURL 单机承诺函
	OfflineAppLetterURL string `json:"offline_app_letter_url,omitempty"`
}

// AppSource 应用创建者信息
type AppSource struct {
	// AccountID 应用创建账号id
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountName 应用创建账号名称
	AccountName string `json:"account_name,omitempty"`
}

// AppPrivacyInfo 隐私声明数据
type AppPrivacyInfo struct {
	// PrivacyID 隐私声明ID
	PrivacyID uint64 `json:"privacy_id,omitempty"`
	// URL 隐私声明链接
	URL string `json:"url,omitempty"`
}

// PackageInfo 应用包数据
type PackageInfo struct {
	// PackageID 应用包ID
	PackageID uint64 `json:"package_id,omitempty"`
	// AppName 应用名称：可不填，如果不填则默认使用上传 APK 时解析出的应用名称。
	AppName string `json:"app_name,omitempty"`
	// BlobStoreKey 应用存储 Key，上传应用 APK 时返回。
	BlobStoreKey string `json:"blob_store_key,omitempty"`
	// SensitivePermissionDesc 敏感权限用途
	SensitivePermissionDesc string `json:"sensitive_permission_desc,omitempty"`
}
