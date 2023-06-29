package target_v2

type Target struct {
	Region               []int64          `json:"region,omitempty"`            // 地域，传值为 [] 表示不限，传递上一级 id 时，children id 可以不传，不允许同时传 parent id 和 children id，地域信息可通过 /region/list 接口获取
	DistrictIDs          []int64          `json:"district_ids,omitempty"`      // 商圈定向，与 region 字段不能同时传、白名单控制，最多选 100 个，可以通过 /rest/openapi/v1/region/district/list 接口获取商圈信息
	BehaviorInterest     BehaviorInterest `json:"behavior_interest,omitempty"` // 行为兴趣定向
	UserType             int              `json:"user_type,omitempty"`         // 用户类型，0：实时，1：常驻，2：不限
	AgesRange            []int            `json:"ages_range,omitempty"`        // 可选年龄段，18：18-23 岁，24：24-30 岁，31：31-40 岁，41：41-49 岁，50：50-100 岁
	Age                  Age              `json:"age,omitempty"`
	Gender               int              `json:"gender,omitempty"`                 // 性别，1：女性，2：男性，0 表示不限
	PlatformOS           int              `json:"platform_os,omitempty"`            // 定向的 os 版本，应用安装类计划以 app_id 中的 platform_os 为准，1：Android，2：iOS，0 表示不限
	AndroidOSV           int              `json:"android_osv,omitempty"`            // Android 版本，3：不限，4：4.x+，5：5.x+，6：6.x+，7：7.x+
	IOSOSV               int              `json:"ios_osv,omitempty"`                // iOS 版本，6：不限，7：7.x+，8：8.x+，9：9.x+，10：10.x+
	Network              int              `json:"network,omitempty"`                // 网络环境，1：WI-FI，2：移动网络，0：表示不限
	FilterConvertedLevel int              `json:"filter_converted_level,omitempty"` // 过滤已转化人群纬度，0：不限，1：广告组，2：广告计划，3：本账户，4：公司主体，5：APP
	DeviceBrand          []int            `json:"device_brand,omitempty"`           // 设备品牌，传值为 [] 表示不限，当 platform 为 iOS 定向时，没有设备品牌定向，1：苹果（platform 字段为空时可选），2：VIVO，3：OPPO
	DevicePrice          []int            `json:"device_price"`                     // 设备价格
	AppInterestIds       []int64          `json:"app_interest_ids"`                 // APP 行为-按分类
	AppIds               []int64          `json:"app_ids"`                          // APP 行为-按 APP 名称
	Population           []int64          `json:"population"`                       // 定向人群包
	PaidAudience         []int64          `json:"paid_audience"`                    // 付费人群包
	ExcludePopulation    []int64          `json:"exclude_population"`               // 排除人群包
	IntelliExtendOption  int              `json:"intelli_extend_option"`            // 智能定向开关
	BehaviorType         int              `json:"behavior_type"`                    // 行为兴趣类型
	Celebrity            Celebrity        `json:"celebrity,omitempty"`              // 快手网红
	MediaSourceType      int              `json:"media_source_type"`                // 媒体包来源
	Media                []int64          `json:"media,omitempty"`                  // 媒体定向包
	ExcludeMedia         []int64          `json:"exclude_media,omitempty"`          // 媒体定向排除包
	SeedPopulation       []int64          `json:"seed_population,omitempty"`        // 种子人群包
	IpType               int              `json:"ip_type,omitempty"`                //地域IP类型 白名单可用，0-默认IP、1-广协IP

}

// Age 自定义年龄段
type Age struct {
	// Min 年龄最小限制
	Min int `json:"min,omitempty"`
	// Max 年龄最大限制
	Max int `json:"max,omitempty"`
}

type Celebrity struct {
	//Behaviors 该字段为平台对应快手网红功能中的行为类型，可多选，但不可不选，不选会导致快手网红定向失效，定义如下 0:关注、1:视频互动、2:直播互动
	Behaviors []int `json:"behaviors"`
	//FansStars 该字段包括平台对应快手网红功能中的网红分类和快手网红两项，数据保证录入顺序，可从【快手网红-网红分类】和【快手网红-搜索快手网红】接口获取对应数据
	FansStars FansStar `json:"fans_stars"`
}

type FansStar struct {
	// ID 如使用该功能则必填此项，该字段有两种含义 1、当是网红分类，即type=1时，该字段传入对应网红分类对应的父ID与当前ID的拼接字符串，
	//如传入"33-177"表示一级网红分类"游戏"下的二级网红分类"沙盒游戏"，如选中的是一级网红分类，则直接传入当前ID如"33"
	//2、当是快手网红，即type=2时，该字段传入对应快手网红的author_id，如传入"1151465119"表示快手网红小脑斧来自N次元
	ID string `json:"id"`
	//Type 如使用该功能则必填此项，该字段为平台对应快手网红功能中的网红分类和快手网红两项，可从【快手网红-网红分类】和【快手网红-搜索快手网红】
	//接口获取对应数据 1:网红分类、2:快手网红
	Type int `json:"type"`
	//Name 如使用该功能则必填此项，该字段有两种含义 1、当是网红分类，即type=1时，该字段传入对应网红分类对应的父name与当前name的拼接字符串，
	//如传入"游戏-沙盒游戏"表示一级网红分类"游戏"下的二级网红分类"沙盒游戏"，如选中的是一级网红分类，则直接传入当前ID如"游戏"，保证传入的值
	//与传入id对应网红标签name（拼接）一致 2、当是快手网红，即type=2时，该字段传入对应快手网红的kwai_id，如传入"小脑斧来自N次元"表示快手
	//网红小脑斧来自N次元，保证传入的值与传入id对应的快手网红的kwai_id保持一致
	Name string `json:"name"`
	//Category 当是快手网红，即type=2时，该字段传值有效，对应当前快手网红的分类，格式为first_label_id,second_label_id，
	//如果second_label_id不存在，则只传入first_label_id，如传入32, 241表示当前快手网红属于影视-影视分类二级网红分类
	Category []int `json:"category"`
}
