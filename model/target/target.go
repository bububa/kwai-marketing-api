package target

// Target 定向数据
type Target struct {
	// Region 地域; 传值为 [] 表示不限，传递上一级 id 时，children id 可以不传，不允许同时传 parent id 和 children id， 地域信息可通过 /region/list 接口获取
	Region []uint64 `json:"region,omitempty"`
	// DistrictIDs 商圈定向; 与region字段不能同时传、白名单控制，最多选100个。 可以通过/rest/openapi/v1/region/district/list接口获取商圈信息
	DistrictIDs []uint64 `json:"district_ids,omitempty"`
	// UserType 用户类型; 0：实时 1：常驻 2：不限
	UserType int `json:"user_type,omitempty"`
	// Age 自定义年龄段 年龄最小限制，大于等于 18 min 和 max 年龄区间最小为 5 岁
	Age *Age `json:"age,omitempty"`
	// AgesRange 固定年龄段; 18：表示18-23岁】【24：表示24-30岁】【31：表示31-40岁】【41：表示41-49岁】【50：表示50-100岁】
	AgesRange []int `json:"ages_range,omitempty"`
	// Gender 性别; 1：女性, 2：男性，0表示不限
	Gender int `json:"gender,omitempty"`
	// PlatformOs 操作系统; 1：Android，2：iOS，0表示不限
	PlatformOs int `json:"platform_os,omitempty"`
	// AndroidOsv Android版本; 3：不限，4：4.x+，5：5.x+，6：6.x+，7：7.x+，8：8.x+，9：9.x+，10：10.x+
	AndroidOsv int `json:"android_osv,omitempty"`
	// IosOsv iOS版本; 6：不限，7：7.x+，8：8.x+，9：9.x+，10：10.x+；
	IosOsv int `json:"ios_osv,omitempty"`
	// Network 网络环境; 1：Wi-Fi，2：移动网络，0：表示不限
	Network int `json:"network,omitempty"`
	// DeviceBrand 设备品牌; 1：OPPO，2：VIVO，3：华为，4：小米，5：荣耀，6：三星，7：金立，8：魅族，9：乐视，10：其他，11：苹果
	DeviceBrand []int `json:"device_brand,omitempty"`
	// DevicePrice 设备价格; 1：1500元以下，2：1501~2000，3：2001~2500，4：2501~3000，5：3001~3500，6：3501~4000，7：4001~4500，8：4501~5000，9： 5001~5500，10：5500元以上
	DevicePrice []int `json:"device_price,omitempty"`
	// BusinessInterestType 商业兴趣类型; 0：不限，1：智能推荐，2：按照兴趣标签；
	BusinessInterestType int `json:"business_interest_type,omitempty"`
	// BusinessInterest 商业兴趣
	BusinessInterest []uint64 `json:"business_interest,omitempty"`
	// FansStar 网红粉丝
	FansStar []int64 `json:"fans_star,omitempty"`
	// InterestVideo 兴趣视频用户
	InterestVideo []int64 `json:"interest_video,omitempty"`
	// AppInterest APP行为-按分类
	AppInterest []uint64 `json:"app_interest,omitempty"`
	// AppInterestIDs APP行为-按分类; id不能重复且必须准确，具体id可通过下方标签接口获取；仅包含安卓数据，若操作系统定向IOS则无效；不能同时选择app_ids（新标签体系字段，替换app_interest，与app_interest同时传递，app_interest字段失效）
	AppInterestIDs []uint64 `json:"app_interest_ids,omitempty"`
	// AppIDs APP行为-按APP名称
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// FilterConvertedLevel 过滤已转化人群纬度; 0(默认)：不限1：广告组2：广告计划3：本账户4：公司主体5：APP
	AppNames []string `json:"app_name,omitempty"`
	// FilterConvertedLevel 过滤已转化人群纬度; 0(默认)：不限1：广告组2：广告计划3：本账户4：公司主体5：APP
	FilterConvertedLevel int `json:"filter_converted_level,omitempty"`
	// Population 人群包定向
	Population []uint64 `json:"population,omitempty"`
	// PaidAudience 付费人群包; 可通过/rest/openapi/v1//dmp/population/list查询 1、付费人群包，人群包 id 不能重复且必须准确 2、人群包状态必须是已推送状态，与定向、排除人群包互斥
	PaidAudience []uint64 `json:"paid_audience,omitempty"`
	// ExcludePopulation 人群包排除; 1、排除人群包，人群包 id 不能重复且必须准确；人群包状态必须是已推送状态 2、如果同时传了 population 和 exclude_population，则表示同时定向排除； 3、population 和 exclude_population 这两个字段不能包含交集的 id
	ExcludePopulation []uint64 `json:"exclude_population,omitempty"`
	// IntelliExtend 智能扩量
	IntelliExtend *IntelliExtend `json:"intelli_extend,omitempty"`
	// BehaviorInterest 行为兴趣定向
	BehaviorInterest *BehaviorInterest `json:"behavior_interest,omitempty"`
	// Media 媒体定向包
	Media []uint64 `json:"media,omitempty"`
	// ExcludeMedia 媒体定向排除包; media和exclude_media只可选其一
	ExcludeMedia []uint64 `json:"exclude_media,omitempty"`
	// MediaSourceType 媒体包来源; 默认为0，0-不限，未指定，1-行业优质流量包，2-广告主自定义
	MediaSourceType int `json:"media_source_type,omitempty"`
}
