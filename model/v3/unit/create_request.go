package unit

import (
	"encoding/json"
)

// CreateRequest 创建广告组APIRequest
type CreateRequest struct {
	//advertiser_id	long	必填	广告主 ID	在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//campaign_id	long	必填	广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	//unit_name	string	必填	广告组名称	长度为 1-100 个字符，同一个计划下的广告组名称不能重复
	UnitName string `json:"unit_name,omitempty"`
	//put_status	int	选填	广告组的投放状态	1：广告组投放中；2：广告组暂停投放；当不传或为 null 时，会默认填充 1
	PutStatus int `json:"put_status,omitempty"`
	//bid_type	int	必填	优化目标出价类型	2：CPC（计划类型为粉丝直播推广、DPA不支持，计划出价为MCB不支持；广告组选择开屏不支持；搜索广告需要添加白名单才可支持）；10：OCPM ；12:最大转化
	BidType int `json:"bid_type,omitempty"`
	//bid	long	选填	出价	bid_type 为 CPC 时该字段必填，单位：厘，不得低于 0.2 元，不得高于 100 元，不得高于组预算
	Bid uint64 `json:"bid,omitempty"`
	//cpa_bid	long	选填	出价	bid_type 是 OCPM 时该字段必填，单位：厘，ocpx_action_type 为 2 时，不得低于 1 元，不得高于 3000 元；ocpx_action_type 为 180 时，不得低于 1 元，不得高于 3000 元，ocpx_action_type 为 53 时，不得低于 5 元，不得高于 3000 元；不得高于组预算
	CpaBid uint64 `json:"cpa_bid,omitempty"`
	//ocpx_action_type	int	选填	优化目标	bid_type 是 OCPM 时该字段必填，需通过接口「rest/openapi/gw/dsp/v1/ocpx/ocpxTypes」获取可以选择的优化目标;；2：行为数（actionBar 点击）；180：激活数；53：表单数；109：电话卡激活；137：量房；190： 付费；191：首日 ROI；324：唤起应用（campaign_type=7+主站+加白）；348：有效线索；383: 授信；384: 完件 394：订单提交；（计划类型获取电商下单，非金牛）；396：注册（联盟暂不支持）；715：微信复制优化目标；（campaign_type=5&bid_type=10）；716：多转化事件(campaign_type=5&bid_type=10&使用魔力建站&线索转化类组件>=2 种)；717：广告观看次数(campaign_type=2&bid_type=10&app与转化追踪工具联调过或者该广告主的在相应事件的回传数>=1）；731：广告观看 5 次；732：广告观看 10 次；733：广告观看 20 次；773：关键行为；774：7 日 ROI；72：小店通商品和主页推广，转化目标“涨粉”；739:7 日付费次数;392：小店通商品推广，优化目标“转化数”，转化目标“商品访问”；395：小店通商品和主页推广，转化目标“订单支付”，62：直播观看数,（小店直播推广&商品推广支持，计划 type=14、13）;192：直播推广 ROI；(小店直播推广&商品推广支持，计划 type=14、13）;634：预约表单；635：预约点击跳转。需通过接口「获取可选的深度转化类型」，接口 is_order_paied 字段返回值为 1 才可以使用; 810:激活付费;346:自然日次日留存
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
	//deep_conversion_type	int	选填	深度转化目标	通过接口「/rest/openapi/gw/dsp/v1/ocpx/deepTypes」获取可以选择深度转化目标；3:付费，7:次日留存，10:完件, 11:授信；13:添加购物车；14:提交订单；15:购买；44：有效线索；92：付费 roi；181：激活后24H次日留存。
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	//enhance_conversion_type	int	选填	增强目标	8：7日留存
	EnhanceConversionType int `json:"enhance_conversion_type,omitempty"`
	//roi_ratio	double	选填	付费 ROI 系数	优化目标为「首日 ROI、7日ROI出价」时必填：ROI 系数取值范围 ( 0,100 ] 最多支持到三位小数（如：0.066）
	RoiRatio float64 `json:"roi_ratio,omitempty"`
	//deep_conversion_bid	long	选填	深度转化目标出价	单位：厘，仅当 deep_conversion_type 可用且账户满足白名单时选填，出价需大于 cpa_bid，小于 min{组预算，3000 元}，不得高于组预算，不支持从 0 改为其他值，也不支持从其他值改为 0
	DeepConversionBid uint64 `json:"deep_conversion_bid,omitempty"`
	//scene_id	int[]	必填	资源位置	1：优选广告位（当账户在联盟优选白名单中且 campaign_type=2/3/5/7 时为主站&联盟优选广告位 ）；5：联盟广告；（对应DSP后台unit层级选择联盟广告位）；6：上下滑大屏广告（计划 campaignType=16 粉丝直播推广时，仅支持填写 6）；10：联盟场景;当账户在联盟优选白名单中且 campaign_type = 2/3/5/7，可传 10；当账户不在联盟优选白名单中或 campaign_type 不等于 2/3/5/7 时，不可传 10；（对应DSP后台创意层级选择联盟广告位）；24:激励视频（需要加白后可选）；27：开屏广告位（白名单功能）；39：搜索广告
	SceneId []int `json:"scene_id,omitempty"`
	//unit_type	int	必填	创意制作方式	3：DPA 自定义创意 4: 常规自定义创意；5：程序化创意 7：程序化创意 2.0
	UnitType int `json:"unit_type,omitempty"`
	//begin_time	string	必填	投放开始时间	格式为 yyyy-MM-dd，需大于等于当前时间
	BeginTime string `json:"begin_time,omitempty"`
	//end_time	string	选填	投放结束时间	格式为 yyyy-MM-dd，不传值表示从今天开始长期投放，传值表示设置开始时间和结束时间，且投放结束时间需大于等于投放开始时间，开屏如果选择按开始时间和结束时间投放，时间(end_time - begin_time)必须超过三天
	EndTime string `json:"end_time,omitempty"`
	//schedule_time	string	选填	投放时间段	格式为 24*7 位字符串，且都为 0 和 1，以一个小时为最小粒度，从周一零点开始至周日 24 点结束；0 为不投放，1 为投放，不传/全传 1/全传 0 视为全时段投放
	ScheduleTime string `json:"schedule_time,omitempty"`
	//day_budget	long	选填	单日预算	单位：厘，指定 0 表示预算不限，默认为 0；每天不小于 100 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该广告组当日花费的 120%，与 day_budget_schedule 不能同时传，均不能低于组出价
	DayBudget uint64 `json:"day_budget,omitempty"`
	//day_budget_schedule	long[]	必填	分日预算	单位：厘，指定 0 表示预算不限，默认为 0；每天不小于 100 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该计划当日花费的 120%，与 day_budget 不能同时传，均不能低于该计划下任一广告组出价，优先级高于day_budget
	DayBudgetSchedule []uint64 `json:"day_budget_schedule,omitempty"`
	//convert_id	int	选填	转化目标 ID	可通接口【/rest/openapi/v1/tool/convert/list】获得，不同计划类型需要对应各自的转化目标类型：提升应用安装(campaign_type=2) - 安卓：convert_type:3、7 / IOS：convert_type:7；推广品牌活动(campaign_type=4) / 收集销售线索(campaign_type=5)：convert_type:1、2 收集销售线索营销目标下，建站落地页程序化落地页不支持convertId，自有落地页仅支持convert_type=1的场景
	ConvertId int `json:"convert_id,omitempty"`
	//url_type	int	选填	url 类型	当计划类型为获取电商下单（campaign.type=3）时必填，其它场景下无效；3 - 手淘落地页（深度唤醒手机淘宝，scene_id为5时联盟场景不可用）；4 - 联盟商选（scene_id为5时联盟场景可用）
	UrlType int `json:"url_type,omitempty"`
	//web_uri_type	int	选填	url 类型	当计划类型为收集销售线索、提升应用活跃（campaign.type = 5/7 ）且使用建站落地页时必填；1：广告主自有链接，默认填充值； 2：建站落地页，此时siteId_id 需要填充；3：程序化落地页，此时group_id需要填充（仅收集营销线索下可用） ；4：微信小程序调起链接（仅收集营销线索下且scene_id包含5可用）；
	WebUriType int `json:"web_uri_type,omitempty"`
	//url 因为跟方法名重复所以结构体中使用WebUrl	string	选填	投放链接	长度不超过 1000 字符；计划类型是获取电商下单（campaign.type=3)，url_type = 3 时表示手淘落地页，url_type = 4 时表示联盟商选落地页ID，必须为数字；计划类型是 4（推广品牌活动）：落地页 URL；web_uri_type = 1 时表示客户自有链接必填；web_uri_type = 4 时表示小程序的落地页；「房地产」「家装建材」「招商加盟」三个二级行业【收集销售线索】目标下隐藏客户自有链接填写入口。
	WebUrl string `json:"url,omitempty"`
	//site_id	long	选填	建站ID/下载中间页ID	当 web_uri_type = 2 时表示建站 ID，必须为数字，通过「/rest/openapi/v2/lp/page/list」 获取。计划类型是2（提升应用安装）且关联应用为安卓时，表示安卓下载中间页ID，通过「/rest/openapi/v2/lp/page/list」 获取 "view_comps = 7" 的建站ID。
	SiteId uint64 `json:"site_id,omitempty"`
	//group_id	long	选填	程序化落地页ID	web_uri_type = 3 时表示程序化落地页ID，必须为数字，通过「/rest/openapi/gw/magicsite/v1/group/list」获取；
	GroupId uint64 `json:"group_id,omitempty"`
	//schema_uri	string	必填	调起链接	提升应用活跃营销目标的调起链接；应用推广下在已经安装 app 的用户手机上可以拉起 app（需要运营加白），开屏广告如果营销目标是应用活跃，调起链接必须在品牌开屏白名单中；当 web_uri_type = 4 时，该字段必填，表示带归因参数的小程序启动页面链接，当campaign.type=5时，白名单用户可填。campaign.type=9 时，若该字段非空，使用该链接唤起应用，否者使用商品上的链接直达商品
	SchemaUri string `json:"schema_uri,omitempty"`
	//schema_id	String	选填	微信小程序ID	收集营销线索下，web_uri_type = 4时生效，表示微信小程序ID
	SchemaId string `json:"schema_id,omitempty"`
	//app_id	long	必填	应用 ID	当计划类型为 提升应用活跃/提升应用安装/商品库推广（campaign.type = 2/7/9） 时必填，可通过应用列表接口「/rest/openapi/v2/file/ad/app/list」获取应用 ID；当计划类型为粉丝直播推广（campaign.type = 16），且组件类型为小铃铛关联了应用时必填；
	AppId uint64 `json:"app_id,omitempty"`
	//app_download_type	int	选填	应用下载方式	0 - 直接下载；1 - 落地页下载；仅在提升应用活跃（默认是直接下载）、提升应用安装（默认是落地页下载）下白名单生效。
	AppDownloadType int `json:"app_download_type,omitempty"`
	//use_app_market	int	选填	优先从系统应用商店下载	0：不从应用商店下载；1：优先从应用商店下载。【填充为1的前提：提升应用活跃、提升应用安装计划类型且选择的应用为安卓。收集营销线索计划类型且选择的建站落地页关联了安卓APP】
	UseAppMarket int `json:"use_app_market,omitempty"`
	//app_store	string[]	选填	应用商店列表	huawei：华为； oppo：OPPO；vivo：VIVO；xiaomi：小米；meizu：魅族；smartisan：锤子 。【仅当use_app_market=1时生效】
	AppStore []string `json:"app_store,omitempty"`
	//show_mode	int	必填	创意展现方式	1 - 轮播；2 - 优选
	ShowMode int `json:"show_mode,omitempty"`
	//site_type	int	加白	预约广告	1 - IOS预约；2 - ANDROID 预约；【仅在计划类型为收集营销线索，web_uri_type = 2 时生效】
	SiteType int `json:"site_type,omitempty"`
	//smart_cover	boolean	unit_type=7 选填	程序化创意 2.0 智能抽帧	是否开启智能抽帧
	SmartCover bool `json:"smart_cover,omitempty"`
	//asset_mining	boolean	unit_type=7 选填	程序化创意 2.0 素材挖掘	是否开启历史素材挖掘
	AssetMining bool `json:"asset_mining,omitempty"`
	//consult_id	long	选填:从工具-获取可选咨询组件	咨询组件 id	1、仅可被用于线索类计划下的 unit；2、仅当落地页使用了建站落地页或者建站程序化落地页时可使用；3、注意本字段不可被更新；4、本属性不可与附加表单组件(component_id)同时使用；通过「/rest/openapi/v2/lp/consult/list」接口获取
	ConsultId uint64 `json:"consult_id,omitempty"`
	//adv_card_option	int	选填	高级创意开关	0：关闭 1:开启
	AdvCardOption int `json:"adv_card_option,omitempty"`
	//adv_card_list	long[]	选填	绑定卡片 id	card_type=100 为 1 个；card_type=101 为 3 个；card_type=102 为 3 个；card_type=103 为 1 个；通过接口「/rest/openapi/v1/asset/adv_card/list」获取
	AdvCardList []uint64 `json:"adv_card_list,omitempty"`
	//card_type	int	选填	卡片类型	100:图片卡片 101:多利益卡-图文 102：多利益卡-多标签 103：电商促销样式 104：快捷评论卡 200：推广位
	CardType int `json:"card_type,omitempty"`
	//playable_id	Long	选填	试玩 ID	可选字段，开启试玩时存在，否则不存在。当用户上传试玩素材成功时返回，用于之后对于广告组中试玩创意的编辑操作。通过接口「 /rest/openapi/gw/dsp/v1/playable/list 」 获取
	PlayableId uint64 `json:"playable_id,omitempty"`
	//play_button	String	选填	试玩按钮文字内容	开启试玩时存在，否则不存在，示例按钮内容如下：1：立即试玩；2：试玩一下；3：立即体验；4：免装试玩；5：免装体验。启用试玩时：默认“立即试玩”，未启用试玩时：内容为空。通过接口「/rest/openapi/gw/dsp/v1/playable/play_buttons」获取
	PlayButton string `json:"play_button,omitempty"`
	//dpa_unit_param	struct	选填	DPA 相关商品信息	当计划类型为商品库推广时必填。
	DpaUnitParam *DpaUnitParam `json:"dpa_unit_param,omitempty"`
	//jingle_bell_id	long	选填	小铃铛组件id	计划 campaignType=16 粉丝直播推广时必填写
	JingleBellId uint64 `json:"jingle_bell_id,omitempty"`
	//live_user_id	long	选填	主播id	计划 campaignType=16 粉丝直播推广时必填写
	LiveUserId uint64 `json:"live_user_id,omitempty"`
	//conversion_type	int	选填	转化途径	计划 campaignType=16 粉丝直播推广时必填写 6
	ConversionType int `json:"conversion_type,omitempty"`
	//extend_search	int	选填	智能扩词开启状态	0：关闭，1：开启，不填默认为0
	ExtendSearch int `json:"extend_search,omitempty"`
	//custom_mini_app_data	struct	选填	推广小程序营销目标小程序信息	计划 campaignType=19 推广快手小程序时必填，具体见下方表格。 具体见下方表格
	CustomMiniAppData *CustomMiniAppData `json:"custom_mini_app_data,omitempty"`
	//target	struct	必填	定向数据	具体见下方表格
	Target *Target `json:"target,omitempty"`
	//template_id	long	选填	定向模板 id	此字段关联到定向信息后，target 字段失效。会将关联的定向信息填充到广告组中。
	TemplateId uint64 `json:"template_id,omitempty"`
	//outer_loop_native	int	选填	是否开启原生	0关闭，1开启，不填默认为0仅在计划 campaignType=2提升应用安装、5收集销售线索、7提升应用活跃、19推广快手小程序时，可开启原生投放。
	OuterLoopNative int `json:"outer_loop_native,omitempty"`
	//kol_user_type	int	选填	原生达人用户类型	2服务号原生，3聚星达人原生，当outer_loop_native为1时此项必填
	KolUserType int `json:"kol_user_type,omitempty"`
	//quick_search	int	选填	搜索快投开关	0：不开启（默认填充）；1：开启搜索快投；
	QuickSearch int `json:"quick_search,omitempty"`
	//target_explore	int	选填	搜索人群探索	0：不开启；1：开启；开启后系统将基于用户搜索行为探索定向人群以外的更多优质人群，仅在 quick_search = 1时生效；
	TargetExplore int `json:"target_explore,omitempty"`
	//package_id	long	必填	新版应用中心，应用包ID	使用新版应用中心投放，该字段必填；老版本应用非必填
	PackageId uint64 `json:"package_id,omitempty"`
}
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
type DpaUnitParam struct {
	//dpa_unit_sub_type	int	选填	商品广告类型	商品广告必填：1-DPA，2-SDPA，3-动态商品卡（2/3在提升应用安装、提升应用活跃、收集营销线索、小程序推广营销目标下可用）
	DpaUnitSubType int `json:"dpa_unit_sub_type,omitempty"`
	//library_id	long	选填	商品库ID	商品广告必填
	LibraryId int64 `json:"library_id,omitempty"`
	//product_id	string	选填	快手商品ID	SDPA必填
	ProductId string `json:"product_id,omitempty"`
	//outer_id	string	选填	外部商品ID	SDPA必填
	OuterId string `json:"outer_id,omitempty"`
	//click_url	string	选填	点击跳转监测链接	DPA可用（scene_id = 5 须在 creative 填写）
	ClickUrl string `json:"click_url,omitempty"`
	//actionbar_click_url	string	选填	Actionbar点击跳转监测链接	DPA可用（scene_id = 5 须在 creative 填写）
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// impression_url	string	选填	曝光监测链接	DPA可用（scene_id = 5 须在 creative 填写）
	ImpressionUrl string `json:"impression_url,omitempty"`
	// dpa_outer_ids String	选填	DPA 商品ID集合	DPA可用（仅scene_id = 5支持），与dpa_category_ids 二选一， 两者并存优先dpa_outer_ids
	DpaOuterIds string `json:"dpa_outer_ids,omitempty"`
	// dpa_category_ids String	选填	DPA商品类目集合	DPA可用。类目 ID 通过商品库类目信息接口获取。类目之间用 - 分隔，如："一级类目ID-二级类目ID-三级类目ID" 或 "一级类目ID"。与 dpa_outer_ids 二选一，两者皆空则全类目投放
	DpaCategoryIds string `json:"dpa_category_ids,omitempty"`
}

type CustomMiniAppData struct {
	//mini_app_id_platform	string	选填	小程序APPID	计划 campaignType=19 推广快手小程序时必填，不能超过30个字符。 注：当账户加了小游戏推广白名单时，不需要填写。
	MiniAppIdPlatform string `json:"mini_app_id_platform,omitempty"`
	//bootstrap_page	string	选填	小程序启动页面	跳转小程序时将直接进入该页面，计划 campaignType=19 小程序推广时必填。例：page/app/index。注：当账户加了小游戏推广白名单时，必填。
	BootstrapPage string `json:"bootstrap_page,omitempty"`
	//bootstrap_params	string	选填	小程序启动参数	启动参数可以追踪页面来源，计划 campaignType=19 小程序推广时选填。例：name=lilei&age=18。注：当账户加了小游戏推广白名单时，不需要填写。
	BootstrapParams string `json:"bootstrap_params,omitempty"`
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

// Url implement PostRequest interface
func (r CreateRequest) Url() string {
	return "/gw/dsp/unit/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
