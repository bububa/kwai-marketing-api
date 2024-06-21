package subpkg

type ListResponse struct {
	CurrentPage int          `json:"current_page"`
	PageSize    int          `json:"page_size"`
	TotalCount  int          `json:"total_count"`
	List        []SubpkgItem `json:"list"`
}

type SubpkgItem struct {
	//account_id	Long		账号ID
	AccountId int64 `json:"account_id"`
	//app_detail_img	String		应用详情图片
	AppDetailImg string `json:"app_detail_img"`
	//app_icon_url	String		应用图标链接
	AppIconUrl string `json:"app_icon_url"`
	//app_id	Long		应用ID
	AppId int64 `json:"app_id"`
	//app_privacy_url	String		应用隐私政策链接
	AppPrivacyUrl string `json:"app_privacy_url"`
	//channel_id String 渠道号(分包号)
	ChannelId string `json:"channel_id"`
	//ios_app_id	String		解析出的iosAppID
	IosAppId string `json:"ios_app_id"`
	//offline_app_stores	String		下架的应用商店	"huawei","oppo","vivo","xiaomi","meizu","smartisan"
	OfflineAppStores string `json:"offline_app_stores"`
	//package_id	Long		应用包ID
	PackageId int64 `json:"package_id"`
	//package_name	String		应用包名
	PackageName string `json:"package_name"`
	//package_size	Long		应用包大小
	PackageSize int64 `json:"package_size"`
	//permission_information	int[]		权限信息ID列表
	PermissionInformation []int `json:"permission_information"`
	//platform	String		android或ios
	Platform string `json:"platform"`
	//real_app_name	String		应用名称
	RealAppName string `json:"real_app_name"`
	//real_app_version	String		应用版本信息
	RealAppVersion string `json:"real_app_version"`
	//source_type	Integer		应用来源	1-我创建的 2-共享给我的
	SourceType int `json:"source_type"`
	//update_time	Long		更新时间	单位：毫秒
	UpdateTime int64 `json:"update_time"`
	//url	String		应用下载地址
	Url string `json:"url"`
	//use_sdk	Integer		是否接入快手广告监测SDK	0-未接入，1-已接入
	UseSdk int `json:"use_sdk"`
	//version_code
	VersionCode int `json:"version_code"`
}
