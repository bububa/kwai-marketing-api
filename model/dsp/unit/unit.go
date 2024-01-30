package unit

import "github.com/bububa/kwai-marketing-api/model"

// Unit 广告组
type Unit struct {
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// PutStatus 投放状态（操作结果）
	// 1：投放中；2：暂停 3：删除
	PutStatus int `json:"put_status,omitempty"`
	// Status 广告组状态（优先看看这个状态，计算结果）
	// -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，，11：审核中，12：审核未通过，14：已结束，15：已暂停，17：组超预算，19：未达投放时间，20：有效，22：不在投放时段
	Status int `json:"status,omitempty"`
	// ReviewDetail 审核拒绝理由
	ReviewDetail string `json:"review_detail,omitempty"`
	// StudyStatus 学习期
	// 1:学习中,2:学习成功,3:学习失败
	StudyStatus int `json:"study_status,omitempty"`
	// CompensateStatus 赔付状态
	// 0=不需要赔付(不需要展示赔付标志)，1=成本保障生效中，2=成本保证确认中，3=已赔付完成，4=已失效
	CompensateStatus int `json:"compensate_status,omitempty"`
	// BidType 出价类型
	// 1：CPM，2：CPC，6：OCPC(使用 OCPC 代表 OCPX)，10：OCPM，20：eCPC
	BidType int `json:"bid_type,omitempty"`
	// Bid 出价
	// 单位：厘
	Bid int64 `json:"bid,omitempty"`
	// CpaBid OCPC 出价
	// 单位：厘
	CpaBid int64 `json:"cpa_bid,omitempty"`
	// OcpxActionType 优化目标
	// 0：未知，2：点击转化链接，10：曝光，11：点击，31：下载完成，53：提交线索，109：电话卡激活，137：量房，180：激活，190: 付费，191：首日 ROI，348：有效线索，383: 授信，384: 完件 715：微信复制;739:7 日付费次数;774:7 日ROI;810:激活付费
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
	// DeepConversionType 深度转化类型
	// 3:付费，7:次日留存，10:完件, 11:授信；13:添加购物车；14:提交订单；15:购买；44：有效线索；92：付费 roi；181：激活后24H次日留存；0：无；
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// DeepConversionBid 深度转化目标出价
	// 白名单功能，单位：厘
	DeepConversionBid int64 `json:"deep_conversion_bid,omitempty"`
	// EnhanceConversionType 增强目标
	EnhanceConversionType int `json:"enhance_conversion_type,omitempty"`
	// RoiRatio 付费 ROI 系数
	// 优化目标为「首日 ROI」时必填：ROI 系数取值范围 ( 0,100 ] 最多支持到三位小数（如：0.066）
	RoiRatio float64 `json:"roi_ratio,omitempty"`
	// SceneID 广告位
	// 1：优选广告位；2：按场景选择广告位-信息流广告（旧广告位，包含上下滑大屏广告）6：上下滑大屏广告；7：双列信息流广告（不包含上下滑大屏广告）24：激励视频；11：快看点场景
	SceneID []int `json:"scene_id,omitempty"`
	// UnitType 创意制作方式
	// 3：DPA 自定义创意（仅在计划类型为商品库推广下使用） 4: 常规自定义创意； 7：程序化创意 2.0
	UnitType int `json:"unit_type,omitempty"`
	// BeginTime 投放开始时间
	// 格式：yyyy-MM-dd
	BeginTime string `json:"begin_time,omitempty"`
	// EndTime 投放结束时间
	// 格式：yyyy-MM-dd,排期不限为 null
	EndTime *string `json:"end_time,omitempty"`
	// Schedule 投放时段
	// 不投放的时段为 null，详请见下方，历史字段，即将废弃
	Schedule *string `json:"schedule,omitempty"`
	// ScheduleTime 投放时段
	// 24*7 的字符串，0 为不投放，1 为投放，例如：0010000000001....0000
	ScheduleTime string `json:"schedule_time,omitempty"`
	// DayBudget 单日预算
	// 单位：厘
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算
	// 单位：厘，单日预算和分日预算同时存在时，以分日预算为准，优先级高于day_budget
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
	// ConvertID 转化目标
	ConvertID int `json:"convert_id,omitempty"`
	// URLType url 类型
	// 当计划类型为 3 时（获取电商下单）时有返回。1：淘宝商品短链 2：淘宝商品 itemID
	URLType int `json:"url_type,omitempty"`
	// WebURLType url 类型
	// 当计划类型为 5（收集销售线索）&使用建站时有返回：需使用魔力建站 不传默认 1，2：落地页
	WebURLType int `json:"web_url_type,omitempty"`
	// URL 落地页链接
	// 计划类型是 2（提升应用安装）：返回应用下载地址；计划类型是 3（获取电商下单）：根据 url_type 返回相应信息；计划类型是 4（推广品牌活动）：返回落地页 url 计划类型是 5（收集销售线索）：返回落地页 url 计划类型是 5（收集销售线索）：建站 idl
	URL string `json:"url,omitempty"`
	// SiteID 建站ID/下载中间页ID
	// 当 web_uri_type = 2 时表示建站 ID，必须为数字，通过「/rest/openapi/v2/lp/page/list」 获取。计划类型是2（提升应用安装）且关联应用为安卓时，表示安卓下载中间页ID，通过「/rest/openapi/v2/lp/page/list」 获取 "view_comps = 7" 的建站ID。
	SiteID uint64 `json:"site_id,omitempty"`
	// PageGroupDetail 程序化落地页信息
	// 广告组ID绑定的程序化落地页组信息
	PageGroupDetail *PageGroupDetail `json:"page_group_detail,omitempty"`
	// SchemaURL 调起链接
	// 提升应用活跃营销目标的调起链接
	SchemaURL string `json:"schema_url,omitempty"`
	// SchemaID 微信小程序外部调起链接
	// 目前只有收集营销线索计划下的联盟广告位该字段才有效
	SchemaID string `json:"schema_id,omitempty"`
	// AppID 应用 ID
	AppID uint64 `json:"app_id,omitempty"`
	// AppDownloadType 应用下载方式
	// 0 - 直接下载；1 - 落地页下载；仅在提升应用活跃（默认是直接下载）、提升应用安装（默认是落地页下载）下白名单生效。
	AppDownloadType int `json:"app_download_type,omitempty"`
	// DownloadPageURL 自定义落地页URL
	// 选择落地页下载方式时的自定义落地页URL，仅在提升应用活跃和提升应用安装下生效，site_id 优先级高于此字段，白名单功能
	DownloadPageURL string `json:"download_page_url,omitempty"`
	// UseAppMarket 优先从系统应用商店下载
	// 0：不从应用商店下载；1：优先从应用商店下载。【填充为1的前提：提升应用活跃、提升应用安装计划类型且选择的应用为安卓。收集营销线索计划类型且选择的建站落地页关联了安卓APP】
	UseAppMarcket int `json:"use_app_market,omitempty"`
	// AppStore 应用商店列表
	// huawei：华为； oppo：OPPO；vivo：VIVO；xiaomi：小米；meizu：魅族；smartisan：锤子 。【仅当use_app_market=1时生效】
	AppStore []string `json:"app_store,omitempty"`
	// DiverseData 应用信息
	DiverseData *DiverseData `json:"diverse_data,omitempty"`
	// ShowMode 创意展现方式
	// 1 - 轮播；2 - 优选（默认）
	ShowMode int `json:"show_mode,omitempty"`
	// SiteType 预约广告
	// 1 - IOS预约；2 - ANDROID 预约；【仅在计划类型为收集营销线索，web_uri_type = 2 时生效】
	SiteType int `json:"site_type,omitempty"`
	// SmartCover 程序化创意 2.0 智能抽帧
	// 是否开启智能抽帧
	SmartCover bool `json:"smart_cover,omitempty"`
	// AssetMining 程序化创意 2.0 素材挖掘
	// 是否开启历史素材挖掘
	AssetMining bool `json:"asset_mining,omitempty"`
	// ConsultID 咨询组件
	// 1、仅可被用于线索类计划下的 unit；2、仅当落地页使用了建站落地页或者建站程序化落地页时可使用；3、注意本字段不可被更新；4、本属性不可与附加表单组件(component_id)同时使用；通过「/rest/openapi/v2/lp/consult/list」接口获取
	ConsultID uint64 `json:"consult_id,omitempty"`
	// AdvCardOption 高级创意开关
	// 0：关闭 1:开启
	AdvCardOption int `json:"adv_card_option,omitempty"`
	// AdvCardList 绑定卡片
	// card_type=100 为 1 个；card_type=101 为 3 个；card_type=102 为 3 个；card_type=103 为 1 个；通过接口「/rest/openapi/v1/asset/adv_card/list」获取
	AdvCardList []uint64 `json:"adv_card_list,omitempty"`
	// PlayableID 试玩 ID
	// 可选字段，开启试玩时存在，否则不存在。当用户上传试玩素材成功时返回，用于之后对于广告组中试玩创意的编辑操作。通过接口「 /rest/openapi/gw/dsp/v1/playable/list 」 获取
	PlayableID uint64 `json:"playable_id,omitempty"`
	// PlayButton 试玩按钮文字内容
	// 开启试玩时存在，否则不存在，示例按钮内容如下：1：立即试玩；2：试玩一下；3：立即体验；4：免装试玩；5：免装体验。启用试玩时：默认“立即试玩”，未启用试玩时：内容为空。通过接口「/rest/openapi/gw/dsp/v1/playable/play_buttons」获取
	PlayButton string `json:"play_button,omitempty"`
	// PlayableOrientation 试玩素材的横竖适配
	// 默认值为-1；0:横竖均可；1:竖屏；2:横屏
	PlayableOrientation int `json:"playable_orientation,omitempty"`
	// PlayableFileName 试玩广告的文件名
	PlayableFileName string `json:"playable_file_name,omitempty"`
	// DpaUnitSubType 商品广告类型
	// 商品广告必填：1-DPA，2-SDPA，3-动态商品卡（2/3在提升应用安装、提升应用活跃、收集营销线索、小程序推广营销目标下可用）
	DpaUintSubType int `json:"dpa_unit_sub_type,omitempty"`
	// LibraryID 商品库ID
	// 商品广告必填
	LibraryID uint64 `json:"library_id,omitempty"`
	// OuterID 外部商品ID
	// SDPA必填
	OuterID string `json:"outer_id,omitempty"`
	// DpaOuterIDs DPA 商品ID集合
	// DPA可用（仅scene_id = 5支持），与dpa_category_ids 二选一， 两者并存优先dpa_outer_ids
	DpaOuterIDs []string `json:"dpa_outer_ids,omitempty"`
	// DpaCategoryIDs DPA 商品类目ID集合
	// DPA可用。类目 ID 通过商品库类目信息接口获取。类目之间用 - 分隔，如："一级类目ID-二级类目ID-三级类目ID" 或 "一级类目ID"。与 dpa_outer_ids 二选一，两者皆空则全类目投放
	DpaCategoryIDs []string `json:"dpa_category_ids,omitempty"`
	// DpaDynamicParams 是否开启动态参数
	// 1 表示开启、0表示关闭。（默认关闭）
	DpaDynamicParams int `json:"dpa_dynamic_params,omitempty"`
	// DpaDynamicParamsForDp DPA 应用直达链接动态参数值
	// 长度不超过 100 字符
	DpaDynamicParamsForDp string `json:"dpa_dynamic_params_for_dp,omitempty"`
	// DpaDynamicParamsForURL DPA 落地页链接动态参数值
	// 长度不超过 100 字符
	DpaDynamicParamsForURL string `json:"dpa_dynamic_params_for_url,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ProductPrice 商品价格, 单位：元
	ProductPrice float64 `json:"product_price,omitempty"`
	// ProductImage 商品主图
	ProductImage string `json:"product_image,omitempty"`
	// JingleBellID 小铃铛组件id
	// 计划 campaignType=16 粉丝直播推广时必填写
	JingleBellID uint64 `json:"jingle_bell_id,omitempty"`
	// LiveUserID 主播id
	// 计划 campaignType=16 粉丝直播推广时必填写, 计划 campaignType=30 短剧推广时，必填，值为短剧作者ID
	LiveUserID uint64 `json:"live_user_id,omitempty"`
	// ExtendSearch 智能扩词开启状态
	// false：关闭，true：开启，不填默认为false
	ExtendSearch bool `json:"extend_search,omitempty"`
	// CustomMiniAppData 推广小程序营销目标小程序信息
	// 计划 campaignType=19 推广快手小程序时必填，具体见下方表格。 具体见下方表格
	CustomMiniAppData *CustomMiniAppData `json:"custom_mini_app_data,omitempty"`
	// Target 定向数据
	Target *Target `json:"target,omitempty"`
	// TemplateID 定向模板 id
	// 此字段关联到定向信息后，target 字段失效。会将关联的定向信息填充到广告组中。
	TemplateID uint64 `json:"template_id,omitempty"`
	// OuterLoopNative 是否开启原生
	// 1开启，0关闭 不填则默认为0，投放快手信息流广告时（ad_type=0或默认不填，scene_id包含1优选广告位、6上下滑大屏广告、7双列信息流广告），当campaignType=2提升应用安装、5收集销售线索、7提升应用活跃、19推广快手小程序、30快手号-短剧推广时，可开启原生投放。注：投放快手信息流广告且升级白名单账户必选 outer_loop_native = 1，否则会报错
	OuterLoopNative int `json:"outer_loop_native,omitempty"`
	// QuickSearch 搜索快投开关
	// 0：不开启（默认填充）；1：开启搜索快投；
	QuickSearch int `json:"quick_search,omitempty"`
	// TargetExplore 搜索人群探索
	// 0：不开启；1：开启；开启后系统将基于用户搜索行为探索定向人群以外的更多优质人群，仅在 quick_search = 1时生效；
	TargetExplore int `json:"target_explore,omitempty"`
	// NegativeWordParam 搜索广告否词
	// 设置否词必填
	NegativeWordParam *NegativeWordParam `json:"negative_word_param,omitempty"`
	// BackflowForcast 回流预估信息
	// 回流预估信息
	BackflowForcast *BackflowForcast `json:"backflow_forcast,omitempty"`
	// AdType 广告计划类型
	// 0:信息流，1:搜索
	AdType int `json:"ad_type,omitempty"`
	// UnitSource 广告组来源
	// 0:常规（非托管）、1:托管
	UnitSource int `json:"unit_source,omitempty"`
	// UpdateTime 最后修改时间
	// 格式样例："2019-06-11 15:17:25"
	UpdateTime string `json:"update_time,omitempty"`
	// PackageID 新版应用中心，应用包ID
	// 使用新版应用中心投放，该字段必填；老版本应用非必填 获取母包接口文档 获取分包接口文档
	PackageID uint64 `json:"package_id,omitempty"`
	// SeriesID 短剧id
	// 当计划 campaignType=30 短剧推广时，必填
	SeriesID uint64 `json:"series_id,omitempty"`
	// EpisodeID 短剧下剧集id
	// 当计划 campaignType=30 短剧推广时，必填
	EpisodeID uint64 `json:"episode_id,omitempty"`
	// SeriesCardType 是否使用短剧卡片
	// 0关闭，1开启，默认关闭，仅在计划 campaignType=30快手号-短剧推广时支持开启使用短剧卡片
	SeriesCardType int `json:"series_card_type,omitempty"`
	// SeriesCardInfo 短剧卡片信息
	// 开启短剧卡片时，必填，具体见下方表格
	SeriesCardInfo *SeriesCardInfo `json:"series_card_info,omitempty"`
	// SeriesPayTemplateID 短剧付费模板id
	// 仅在计划 campaignType=30快手号-短剧推广时支持使用，需要跟付费模式一起使用
	SeriesPayTemplateID uint64 `json:"series_pay_template_id,omitempty"`
	// SeriesPayMode 短剧付费模式
	// 仅在计划 campaignType=30快手号-短剧推广时，需要跟付费模板一起使用，1-打包，2-虚拟币
	SeriesPayMode int `json:"series_pay_mode,omitempty"`
	// ULink Universal Link 链接
	// 仅在计划 campaignType=7 提升应用活跃时使用，输入后IOS将优先调起该链接，不超过2000字符
	ULink string `json:"u_link,omitempty"`
}

// DpaUnitParam DPA 相关商品信息
type DpaUnitParam struct {
	// DpaUnitSubType 商品广告类型
	// 商品广告必填：1-DPA，2-SDPA，3-动态商品卡（2/3在提升应用安装、提升应用活跃、收集营销线索、小程序推广营销目标下可用）
	DpaUintSubType int `json:"dpa_unit_sub_type,omitempty"`
	// LibraryID 商品库ID
	// 商品广告必填
	LibraryID uint64 `json:"library_id,omitempty"`
	// ProductID 快手商品ID
	// SDPA必填
	ProductID string `json:"product_id,omitempty"`
	// OuterID 外部商品ID
	// SDPA必填
	OuterID string `json:"outer_id,omitempty"`
	// ClickURL 点击跳转监测链接
	// DPA可用（scene_id = 5 须在 creative 填写）
	ClickURL string `json:"click_url,omitempty"`
	// ActionbarClickURL Actionbar点击跳转监测链接
	// DPA可用（scene_id = 5 须在 creative 填写）
	ActionbarClickURL string `json:"actionbar_click_url,omitempty"`
	// ImpressionURL 曝光监测链接,DPA可用（scene_id = 5不支持）
	// DPA可用
	ImpressionURL string `json:"impression_url,omitempty"`
	// DpaOuterIDs DPA 商品ID集合
	// DPA可用（仅scene_id = 5支持），与dpa_category_ids 二选一， 两者并存优先dpa_outer_ids
	DpaOuterIDs []string `json:"dpa_outer_ids,omitempty"`
	// DpaCategoryIDs DPA 商品类目ID集合
	// DPA可用。类目 ID 通过商品库类目信息接口获取。类目之间用 - 分隔，如："一级类目ID-二级类目ID-三级类目ID" 或 "一级类目ID"。与 dpa_outer_ids 二选一，两者皆空则全类目投放
	DpaCategoryIDs []string `json:"dpa_category_ids,omitempty"`
	// DpaDynamicParams 是否开启动态参数
	// 1 表示开启、0表示关闭。（默认关闭）
	DpaDynamicParams int `json:"dpa_dynamic_params,omitempty"`
	// DpaDynamicParamsForDp DPA 应用直达链接动态参数值
	// 长度不超过 100 字符
	DpaDynamicParamsForDp string `json:"dpa_dynamic_params_for_dp,omitempty"`
	// DpaDynamicParamsForURL DPA 落地页链接动态参数值
	// 长度不超过 100 字符
	DpaDynamicParamsForURL string `json:"dpa_dynamic_params_for_url,omitempty"`
}

// CustomMiniAppData 推广小程序营销目标小程序信息
type CustomMiniAppData struct {
	// MiniAppType 小程序类型
	// 1 表示小程序；2 表示小游戏。(默认为 1)
	MiniAppType int `json:"mini_app_type,omitempty"`
	// MiniAppIDPlatform 小程序APPID
	// 计划 campaignType=19 推广快手小程序时必填，不能超过30个字符。
	MiniAppIDPlatform string `json:"mini_app_id_platform,omitempty"`
	// BootstrapPage 小程序启动页面
	// 当 mini_app_type = 1 时必填，表示跳转小程序时直接进入的页面，当 mini_app_type = 2 时不填。例：page/app/index。
	BootstrapPage string `json:"bootstrap_page,omitempty"`
	// BootstrapParams 小程序启动参数
	// 启动参数可以追踪页面来源
	BootstrapParams string `json:"bootstrap_params,omitempty"`
}

// NegativeWordParam 搜索广告否词
type NegativeWordParam struct {
	// ExactWords 精确否定词
	// 搜索广告新增
	ExactWords []string `json:"exact_words,omitempty"`
	// PhraseWords 短语否定词
	// 搜索广告新增
	PhraseWords []string `json:"phrase_words,omitempty"`
}

// SeriesCardInfo 短剧卡片信息
type SeriesCardInfo struct {
	// PicID 短剧卡片封面
	// 图片比例3:4，仅支持PNG/JPEG/JPG，图片不能大于2M
	PicID uint64 `json:"pic_id,omitempgy"`
	// Label 短剧卡片标签
	// 『专属推荐』、『口碑好剧』、『本月精品』、『本月热播』、『爽文改编』五选二
	Label []string `json:"label,omitempty"`
	// Description 短剧卡片简介
	// 不超过100个字，若选择开启短剧卡片，该字段必填
	Description string `json:"description,omitempty"`
}

// PageGroupDetail 程序化落地页信息
type PageGroupDetail struct {
	// GroupID 程序化落地页组ID
	GroupID uint64 `json:"group_id,omitempty"`
	// GroupName 程序化落地页组名称
	GroupName string `json:"group_name,omitempty"`
	// PageReviewDetail 程序化落地页组下的页面信息
	PageReviewDetail []PageReviewDetail `json:"page_review_detail,omitempty"`
}

// PageReviewDetail 程序化落地页组下的页面信息
type PageReviewDetail struct {
	// PageID 落地页ID
	PageID uint64 `json:"page_id,omitempty"`
	// ReviewStatus 落地页审核状态
	// 1-审核中 2-审核通过 3-审核拒绝
	ReviewStatus model.Int `json:"review_status,omitempty"`
	// ReviewDetail 落地页的审核拒绝理由
	ReviewDetail string `json:"review_detail,omitempty"`
	// URL 落地页URL信息
	URL string `json:"url,omitempty"`
}

// DiverseData 应用信息
type DiverseData struct {
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// AppPackageName 应用包名
	AppPackageName string `json:"app_package_name,omitempty"`
	// DeviceOSType 应用操作系统类型
	// 0：未知，1：ANDROID，2：IO
	DeviceOSType int `json:"device_os_type,omitempty"`
}

// BackflowForcast 回流预估信息
type BackflowForcast struct {
	// BackflowCvLower 回流预估值的下限
	BackflowCvLower int64 `json:"backflow_cv_lower,omitempty"`
	// BackflowCvUpper 回流预估值的上限
	BackflowCvUpper int64 `json:"backflow_cv_upper,omitempty"`
	// BackflowTimestamp 本次回流预估数据的时间戳，13 位毫秒时间戳
	BackflowTimestamp int64 `json:"backflow_timestamp,omitempty"`
	// BackflowPayment 回流预估总金额
	BackflowPayment float64 `json:"backflow_payment,omitempty"`
	// BackflowROI 回流首日预估 ROI
	BackflowROI float64 `json:"backflow_roi,omitempty"`
}
