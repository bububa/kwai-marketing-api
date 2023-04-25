package unit

type Target struct {
	//region	long[]	选填	地域	传值为[]表示不限；传递上一级 ID 时，childrenID 可以不传；不允许同时传 parentID 和 childrenID；地域信息可通过地域接口获取；仅计划的 campaign_type 为 5 时，支持设置三级地域（例：山西-大同-左云，左云是三级地域）
	Region []uint64 `json:"region,omitempty"`
	//district_ids	long[]	选填	商圈定向	与 region 字段不能同时传、白名单控制，最多选 100 个。可以通过/rest/openapi/v1/region/district/list 接口获取商圈信息。该定向仅支持快手站内广告位，不支持联盟广告位。
	DistrictIds []uint64 `json:"district_ids,omitempty"`
	//user_type	int	选填	用户类型	0：实时；1：常驻；2：不限
	UserType int `json:"user_type,omitempty"`
	//age	struct	选填	自定义年龄段	不传值表示不限，传值具体见下方表格；与 ages_range 不能同时传
	Age *Age `json:"age,omitempty"`
	//ages_range	int[]	选填	固定年龄段	与 age 不能同时传；【18：表示 18-23 岁】；【24：表示 24-30 岁】；【31：表示 31-40 岁】；【41：表示 41-49 岁】；【50：表示 50-100 岁】
	AgesRange []int `json:"ages_range,omitempty"`
	//gender	int	选填	性别	1：女性, 2：男性，0 表示不限
	Gender int `json:"gender"`
	//platform_os	int	选填	操作系统	1：Android，2：iOS，0 表示不限；当计划类型为 2（提升应用安装）时该字段可忽略
	PlatformOs int `json:"platform_os,omitempty"`
	//android_osv	int	选填	Android 版本	3：不限，4：4.x+，5：5.x+，6：6.x+，7：7.x+，8：8.x+，9：9.x+，10：10.x+；当计划类型为 2（提升应用安装）时，仅当 app_id 表示的是 Android 应用可使用此字段传递版本
	AndroidOsv int `json:"android_osv,omitempty"`
	//ios_osv	int	选填	iOS 版本	6：不限，7：7.x+，8：8.x+，9：9.x+，10：10.x+；11：11.x+；12：12.x+；13：13.x+；当计划类型为 2（提升应用安装）时，仅当 appId 表示的是 iOS 应用可使用此字段传递版本
	IosOsv int `json:"ios_osv,omitempty"`
	//network	int	选填	网络环境	1：Wi-Fi，2：移动网络，0：表示不限
	Network int `json:"network,omitempty"`
	//device_brand_ids	long[]	选填	设备品牌ID列表	传值为[]表示不限；当 platform为iOS定向时，没有设备品牌定向；1：苹果（platform字段为空时选填）；2：VIVO；3：OPPO；4：华为；5：小米；6：荣耀；7：三星；8：魅族；9：金立；10：HTC；11：联想；12：酷派；13：酷比；14：中兴；15：乐视；16：LG；17：中国移动；18：360；19：百度；20：努比亚；21：一加；22：海信；23：朵唯；24：美图；25：锤子；26：谷歌；27：小辣椒；28：诺基亚；29：康佳
	DeviceBrandIds []uint64 `json:"device_brand_ids,omitempty"`
	//device_price	int[]	选填	设备价格	1：1,500 元以下（近期下线）， 新增：11：1000以下，12：1001~1500，请直接使用新枚举值； 2：1,501~2,000； 3：2,001~2,500；4：2,501~3,000；5、3,001~3,500； 6：3,501~4,000；7：4,001~4,500； 8：4,501~5,000；9：5,001~5,500； 10：5,500 元以上；传值为 [] 表示不限；
	DevicePrice []int `json:"device_price,omitempty"`
	//app_interest_ids	long[]	选填	APP 行为-按分类	id 不能重复且必须准确，具体 id 可通过下方标签接口获取；仅包含安卓数据，若操作系统定向 IOS 则无效；不能同时选择 app_ids（新标签体系字段，替换 app_interest，与 app_interest 同时传递，app_interest 字段失效）。该定向仅支持快手站内广告位，不支持联盟广告位。
	AppInterestIds []uint64 `json:"app_interest_ids,omitempty"`
	//app_ids	long[]	选填	APP 行为-按 APP 名称	id 不能重复且必须准确，具体 id 可通过下方应用接口获取，建议不超过 10 个，否则可能出现报错；仅包含安卓数据，若操作系统定向 IOS 则无效；不能同时选择 app_interest。该定向仅支持快手站内广告位，不支持联盟广告位。
	AppIds []uint64 `json:"app_ids,omitempty"`
	//filter_converted_level	int	选填	过滤已转化人群纬度	搜索广告、联盟广告、小店通不支持。优化目标不支持【封面曝光数】和【封面点击数】非应用下载类推广不支持过滤【APP】纬度。0(默认)：不限；1：广告组 2：广告计划；3：本账户；4：公司主体；5：APP；6:运营自定义产品名
	FilterConvertedLevel int `json:"filter_converted_level,omitempty"`
	//population	long[]	选填	人群包定向	/rest/openapi/v1/dmp/population/list 获取人群包 id。population 不能重复，与 exclude_population 不能有交集，不能传付费人群包，付费人群包标识字段：population_type=7
	Population []uint64 `json:"population,omitempty"`
	//exclude_population	long[]	选填	人群包排除	/rest/openapi/v1/dmp/population/list 获取人群包 id。exclude_population 不能重复，不能传付费人群包，付费人群包标识字段：population_type=7
	ExcludePopulation []uint64 `json:"exclude_population,omitempty"`
	//paid_audience	long[]	选填	付费人群包 id	/rest/openapi/v1/dmp/population/list 获取人群包 id。1、不能重复，只能传付费人群包，且 third_platform_code 要一样，2、如果传了【population 或 exclude_population】，报错；修改时要也要保证【population 或 exclude_population】不能同时传。3、创建时若已经有了【population 或 exclude_population】修改时只传 paid_audience 也会报错，这时需要把 population 或 exclude_population】设为[]才行，反之也一样。总之，【population 或 exclude_population】与【paid_audience】不能同时存在。该定向仅支持快手站内广告位，不支持联盟广告位。
	PaidAudience []uint64 `json:"paid_audience,omitempty"`
	//seed_population	long[]	选填	种子人群包	当账户开通种子人群包功能后方可使用此功能。该定向仅支持快手站内广告位，不支持联盟广告位。
	SeedPopulation []uint64 `json:"seed_population,omitempty"`
	//intelli_extend_option	int	选填	智能定向开关	0：智能定向关闭；1：智能定向开启【计划层级开启智能调控时，只能传1】；当计划层级开启了自动调控功能，不传默认智能定向开启状态，其他清空默认智能定向关闭。
	IntelliExtendOption int `json:"intelli_extend_option,omitempty"`
	//behavior_type	int	选填	行为兴趣类型	0：默认；1：自定义行为意向；2：行为意向系统优选。联盟广告位不支持后者。
	BehaviorType int `json:"behavior_type,omitempty"`
	//behavior_interest	struct	选填	行为兴趣定向	behavior.keyword 、behavior.label、interest.lable 其中一个必传，具体传值下方表格；仅在 behavior_type = 1 时生效。联盟广告位不支持interest。
	BehaviorInterest *BehaviorInterest `json:"behavior_interest,omitempty"`
	//disable_installed_app_switch	int	选填	过滤已安装人群维度	0：过滤（默认），1：不限
	DisableInstalledAppSwitch int `json:"disable_installed_app_switch,omitempty"`
	//filter_time_range	int	选填	用户的转化时间范围	0：30天；1：60天；2:90天。该定向仅支持快手站内广告位，不支持联盟广告位。
	FilterTimeRange int `json:"filter_time_range,omitempty"`
	//celebrity	struct	选填	快手网红	behaviors、fans_star两个字段必填充。该定向仅支持快手站内广告位，不支持联盟广告位。
	Celebrity *Celebrity `json:"celebrity,omitempty"`
	//media	long[]	选填	媒体定向包	若选择使用定向行业优质流量包（media_source_type = 1），此处单选：1 - 超休闲游戏；2 - 重度游戏；4 - 直营电商；5 - 平台电商；若选择使用广告主自定义媒体包（media_source_type = 2），此处多选，填写媒体包 id【仅联盟广告位生效】
	Media []uint64 `json:"media,omitempty"`
	//exclude_media	long[]	选填	媒体定向排除包	media 和 exclude_media 只可选其一，若选择使用排除行业优质流量包（media_source_type = 1），此处单选：1 - 超休闲游戏；2 - 重度游戏；4 - 直营电商；5 - 平台电商；若选择使用广告主自定义媒体包（media_source_type = 2），此处多选，填写媒体包 id；【仅联盟广告位生效】
	ExcludeMedia []uint64 `json:"exclude_media,omitempty"`
	//media_source_type	int	选填	媒体包来源	默认为 0，0-不限，未指定，1-行业优质流量包，2-广告主自定义
	MediaSourceType int `json:"media_source_type,omitempty"`
}

type Age struct {
	// min	int	必填	年龄最小限制	年龄区间最小为 18 岁
	Min int `json:"min"`
	// max	int	必填	年龄最大限制	年龄区间最大为 55 岁，且年龄最大限制须大于等于年龄最小限制
	Max int `json:"max"`
}

type BehaviorInterest struct {
	//behavior	struct	必填	行为定向	behavior 与 interest 同时不传，则清空行为兴趣定向
	Behavior *Behavior `json:"behavior,omitempty"`
	// interest	struct	选填	兴趣定向	behavior 与 interest 同时不传，则清空行为兴趣定向
	Interest *Interest `json:"interest,omitempty"`
}

type Behavior struct {
	//keyword	struct[]	选填	行为定向关键词	根据/rest/openapi/v1/tool/keyword/query 接口获取
	Keyword []*Keyword `json:"keyword,omitempty"`
	//label	string[]	选填	行为定向 类目词	根据/rest/openapi/v1/tool/label/behavior_interest 接口获取。将行为类目 id 从最高层类目 id 开始，以“-”连接起来，假如有一个类目 id 为 80202，父类目 id 为 802，最高层类目 id 为 8，则此时应该写"8-802-80202"；如果想全选最高层类目"8"底下的所有子类目，填"8"
	Label []string `json:"label,omitempty"`
	//time_type	int	必填	在多少天内发生行为的用户	0:7 天；1：15 天；2：30 天；3：90 天；4：180 天
	TimeType int `json:"time_type"`
	//strength_type	int	必填	行为强度	0：不限；1：高强度
	StrengthType int `json:"strength_type"`
	//scene_type	int[]	必填	行为场景	1：社区；2：APP；3:电商；4：推广
	SceneType []int `json:"scene_type"`
}

type Interest struct {
	//label	string[]	选填	兴趣定向类目词	根据/rest/openapi/v1/tool/label/behavior_interest 接口获取。将兴趣类目 id 从最高层类目 id 开始，以“-”连接起来，假如有一个类目 id 为 80202，父类目 id 为 802，最高层类目 id 为 8，则此时应该写"8-802-80202"；如果想全选最高层类目"8"底下的所有子类目，填"8"
	Label []string `json:"label,omitempty"`
	//strength_type	int	必填	兴趣标签强度	0：不限 1：高强度
	StrengthType int `json:"strength_type"`
}

type Keyword struct {
	//id	long	必填	关键词 id	id 与 name 需互相匹配
	Id int64 `json:"id"`
	//name	string	必填	关键词名称	id 与 name 需互相匹配
	Name string `json:"name"`
}

type Celebrity struct {
	//behaviors int[]    必填    行为类型    该字段为平台对应快手网红功能中的行为类型，可多选，但不可不选，不选会导致快手网红定向失效，定义如下 0:关注、1:视频互动、2:直播互动
	Behaviors []int `json:"behaviors"`
	//fans_stars    fans_star[]    必填    网红内容    该字段包括平台对应快手网红功能中的网红分类和快手网红两项，数据保证录入顺序，可从【快手网红-网红分类】和【快手网红-搜索快手网红】接口获取对应数据
	FansStars []FansStar `json:"fans_stars"`
}

type FansStar struct {
	//id    string            如使用该功能则必填此项，该字段有两种含义 1、当是网红分类，即type= 1时，该字段传入对应网红分类对应的父ID与当前ID的拼接字符串，如传入"33-177"表示一级网红分类"游戏"下的二级网红分类"沙盒游戏"，如选中的是一级网红分类，则直接传入当前ID如"33" 2、当是快手网红，即type = 2时，该字段传入对应快手网红的author_id，如传入"1151465119"表示快手网红小脑斧来自N次元
	Id string `json:"id"`
	//type int 如使用该功能则必填此项 该字段为平台对应快手网红功能中的网红分类和快手网红两项，可从【快手网红-网红分类】和【快手网红-搜索快手网红】接口获取对应数据 1:网红分类、2:快手网红
	Type int `json:"type"`
	//name    string            如使用该功能则必填此项，该字段有两种含义 1、当是网红分类，即type = 1时，该字段传入对应网红分类对应的父name与当前name的拼接字符串，如传入"游戏-沙盒游戏"表示一级网红分类"游戏"下的二级网红分类"沙盒游戏"，如选中的是一级网红分类，则直接传入当前ID如"游戏"，保证传入的值与传入id对应网红标签name（拼接）一致 2、当是快手网红，即type = 2时，该字段传入对应快手网红的kwai_id，如传入"小脑斧来自N次元"表示快手网红小脑斧来自N次元，保证传入的值与传入id对应的快手网红的kwai_id保持一致
	Name string `json:"name"`
	//category    int[]            当是快手网红，即type= 2时，该字段传值有效，对应当前快手网红的分类，格式为first_label_id,second_label_id，如果second_label_id不存在，则只传入first_label_id，如传入32, 241表示当前快手网红属于影视-影视分类二级网红分类
	Category []int `json:"category"`
}
