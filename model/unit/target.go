package unit

// Target 定向数据
type Target struct {
	// Region 地域
	Region []int64 `json:"region,omitempty"`
	// DistrictIDs 商圈定向
	DistrictIDs []int64 `json:"district_ids,omitempty"`
	// Age 自定义年龄段
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
	BusinessInterest []int64 `json:"business_interest,omitempty"`
	// FansStar 网红粉丝
	FansStar []int64 `json:"fans_star,omitempty"`
	// InterestVideo 兴趣视频用户
	InterestVideo []int64 `json:"interest_video,omitempty"`
	// AppInterest APP行为-按分类
	AppInterest []int64 `json:"app_interest,omitempty"`
	// AppIDs APP行为-按APP名称
	AppIDs []int64 `json:"app_ids,omitempty"`
	// AppNames app名称; 和app_ids对应，是app_id对应的名称
	AppNames []string `json:"app_name,omitempty"`
	// FilterConvertedLevel 过滤已转化人群纬度; 0(默认)：不限1：广告组2：广告计划3：本账户4：公司主体5：APP
	FilterConvertedLevel int `json:"filter_converted_level,omitempty"`
	// Population 人群包定向
	Population []int64 `json:"population,omitempty"`
	// ExcludePopulation 人群包排除
	ExcludePopulation []int64 `json:"exclude_population,omitempty"`
	// Media 媒体定向包
	Media []int64 `json:"media,omitempty"`
	// ExcludeMedia 媒体定向排除包; media和exclude_media只可选其一
	ExcludeMedia []int64 `json:"exclude_media,omitempty"`
	// MediaSourceType 媒体包来源; 默认为0，0-不限，未指定，1-行业优质流量包，2-广告主自定义
	MediaSourceType int `json:"media_source_type,omitempty"`
	// IntelliExtend 智能扩量
	IntelliExtend *IntelliExtend `json:"intelli_extend,omitempty"`
	// BehaviorInterest 行为兴趣定向
	BehaviorInterest *BehaviorInterest `json:"behavior_interest,omitempty"`
}
