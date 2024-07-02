package unit

import "github.com/bububa/kwai-marketing-api/model"

// UpdateRequest 修改广告组 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	// 长度为 1-100 个字符，同一个计划下的广告组名称不能重复
	UnitName string `json:"unit_name,omitempty"`
	// PutStatus 广告组的投放状态
	// 1：广告组投放中；2：广告组暂停投放；当不传或为 null 时，会默认填充 1
	PutStatus int `json:"put_status,omitempty"`
	// BidType 优化目标出价类型
	// 2：CPC（计划类型为粉丝直播推广、DPA不支持，计划出价为MCB不支持；广告组选择开屏不支持；搜索广告需要添加白名单才可支持）；10：OCPM ；12:MCB最大转化（当计划bid_type=1时，组层级bid_type必须传12）
	BidType int `json:"bid_type,omitempty"`
	// Bid 出价
	// bid_type 为 CPC 时该字段必填，单位：厘，不得低于 0.2 元，不得高于 100 元，不得高于组预算
	Bid int64 `json:"bid,omitempty"`
	// CpaBid 出价
	// bid_type 是 OCPM 时该字段必填，单位：厘，ocpx_action_type 为 2 时，不得低于 0.1 元，不得高于 10 元；ocpx_action_type 为 180 时，不得低于 1 元，不得高于 3000 元，ocpx_action_type 为 53 时，不得低于 5 元，不得高于 3000 元；不得高于组预算
	CpaBid int64 `json:"cpa_bid,omitempty"`
	// DeepConversionBid 深度转化目标出价
	// 单位：厘，仅当 deep_conversion_type 可用且账户满足白名单时选填，出价需大于 cpa_bid，小于 min{组预算，3000 元}，不得高于组预算，不支持从 0 改为其他值，也不支持从其他值改为 0
	DeepConversionBid int64 `json:"deep_conversion_bid,omitempty"`
	// DeepConversionType 深度转化目标
	// 通过接口「/rest/openapi/gw/dsp/v1/ocpx/deepTypes」获取可以选择深度转化目标；3:付费，7:次日留存，10:完件, 11:授信；13:添加购物车；14:提交订单；15:购买；44：有效线索；92：付费 roi；181：激活后24H次日留存。
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// RoiRatio 付费 ROI 系数
	// 优化目标为「首日 ROI、7日ROI出价」时必填：ROI 系数取值范围 ( 0,100 ] 最多支持到三位小数（如：0.066）
	RoiRatio float64 `json:"roi_ratio,omitempty"`
	// SceneID 资源位置
	// 1：优选广告位（当账户在联盟优选白名单中且 campaign_type=2/3/5/7 时为主站&联盟优选广告位 ）；5：联盟广告；（对应DSP后台unit层级选择联盟广告位）；6：上下滑大屏广告（计划 campaignType=16 粉丝直播推广时，仅支持填写 6）；7: 双列信息流广告（加白可用）；10：联盟场景;当账户在联盟优选白名单中且 campaign_type = 2/3/5/7，可传 10；当账户不在联盟优选白名单中或 campaign_type 不等于 2/3/5/7 时，不可传 10；（对应DSP后台创意层级选择联盟广告位）；24:激励视频（需要加白后可选）；27：开屏广告位（白名单功能）；39：搜索广告
	SceneID []int `json:"scene_id,omitempty"`
	// BeginTime 投放开始时间
	// 格式为 yyyy-MM-dd，需大于等于当前时间
	BeginTime string `json:"begin_time,omitempty"`
	// EndTime 投放结束时间
	// 格式为 yyyy-MM-dd，不传值表示从今天开始长期投放，传值表示设置开始时间和结束时间，且投放结束时间需大于等于投放开始时间，开屏如果选择按开始时间和结束时间投放，时间(end_time - begin_time)必须超过三天
	EndTime *string `json:"end_time,omitempty"`
	// ScheduleTime 投放时间段
	// 格式为 24*7 位字符串，且都为 0 和 1，以一个小时为最小粒度，从周一零点开始至周日 24 点结束；0 为不投放，1 为投放，不传/全传 1/全传 0 视为全时段投放
	ScheduleTime string `json:"schedule_time,omitempty"`
	// DayBudget 单日预算
	// 单位：厘，指定 0 表示预算不限，默认为 0；每天不小于 100 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该广告组当日花费的 120% 和修改前预算的较小者，与 day_budget_schedule 不能同时传，均不能低于组出价
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算
	// 单位：厘，指定 0 表示预算不限，默认为 0；每天不小于 100 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该计划当日花费的 120% 和修改前预算的较小者，与 day_budget 不能同时传，均不能低于该计划下任一广告组出价，优先级高于day_budget
	DayBudgetSchedule int64 `json:"day_budget_schedule,omitempty"`
	// URLType url 类型
	// 当计划类型为获取电商下单（campaign.type=3）时必填，其它场景下无效；3 - 手淘落地页（深度唤醒手机淘宝，scene_id为5时联盟场景不可用）
	URLType int `json:"url_type,omitempty"`
	// WebURLType url 类型
	// 当计划类型为收集销售线索、提升应用活跃（campaign.type = 5/7 ）且使用建站落地页时必填；1：广告主自有链接，默认填充值； 2：建站落地页，此时siteId_id 需要填充；3：程序化落地页，此时group_id需要填充（仅收集营销线索下可用） ；4：微信小程序调起链接（仅收集营销线索下且scene_id包含5可用）；
	WebURLType int `json:"web_url_type,omitempty"`
	// URL 投放链接
	// 长度不超过 1000 字符；计划类型是获取电商下单（campaign.type=3)，url_type = 3 时表示手淘落地页；计划类型是 4（推广品牌活动）：落地页 URL；web_uri_type = 1 时表示客户自有链接必填；web_uri_type = 4 时表示小程序的落地页；「房地产」「家装建材」「招商加盟」三个二级行业【收集销售线索】目标下隐藏客户自有链接填写入口。
	URL string `json:"url,omitempty"`
	// SiteID 建站ID/下载中间页ID
	// 当 web_uri_type = 2 时表示建站 ID，必须为数字，通过「/rest/openapi/v2/lp/page/list」 获取。计划类型是2（提升应用安装）且关联应用为安卓时，表示安卓下载中间页ID，通过「/rest/openapi/v2/lp/page/list」 获取 "view_comps = 7" 的建站ID。
	SiteID uint64 `json:"site_id,omitempty"`
	// GroupID 程序化落地页ID
	// web_uri_type = 3 时表示程序化落地页ID，必须为数字，通过「/rest/openapi/gw/magicsite/v1/group/list」获取；
	GroupID uint64 `json:"group_id,omitempty"`
	// SchemaURI 调起链接
	// 提升应用活跃营销目标的调起链接；应用推广下在已经安装 app 的用户手机上可以拉起 app（需要运营加白），开屏广告如果营销目标是应用活跃，调起链接必须在品牌开屏白名单中；当 web_uri_type = 4 时，该字段必填，表示带归因参数的小程序启动页面链接，当campaign.type=5时，白名单用户可填。campaign.type=9 时，若该字段非空，使用该链接唤起应用，否者使用商品上的链接直达商品
	SchemaURI string `json:"schema_uri,omitempty"`
	// SchemaID 微信小程序ID
	// 收集营销线索下，web_uri_type = 4时生效，表示微信小程序ID
	SchemaID string `json:"schema_id,omitempty"`
	// AppID 应用 ID
	// 当计划类型为 提升应用活跃/提升应用安装/商品库推广（campaign.type = 2/7/9） 时必填，可通过应用列表接口「/rest/openapi/v2/file/ad/app/list」获取应用 ID；当计划类型为粉丝直播推广（campaign.type = 16），且组件类型为小铃铛关联了应用时必填；
	AppID uint64 `json:"app_id,omitempty"`
	// AppDownloadType 应用下载方式
	// 0 - 直接下载；1 - 落地页下载；仅在提升应用活跃（默认是直接下载）、提升应用安装（默认是落地页下载）下白名单生效。
	AppDownloadType *int `json:"app_download_type,omitempty"`
	// DownloadPageURL 自定义落地页URL
	// 选择落地页下载方式时的自定义落地页URL，仅在提升应用活跃和提升应用安装下生效，site_id 优先级高于此字段，白名单功能
	DownloadPageURL string `json:"download_page_url,omitempty"`
	// UseAppMarket 优先从系统应用商店下载
	// 0：不从应用商店下载；1：优先从应用商店下载。【填充为1的前提：提升应用活跃、提升应用安装计划类型且选择的应用为安卓。收集营销线索计划类型且选择的建站落地页关联了安卓APP】
	UseAppMarcket int `json:"use_app_market,omitempty"`
	// AppStore 应用商店列表
	// huawei：华为； oppo：OPPO；vivo：VIVO；xiaomi：小米；meizu：魅族；smartisan：锤子 。【仅当use_app_market=1时生效】
	AppStore []string `json:"app_store,omitempty"`
	// ShowMode 创意展现方式
	// 1 - 轮播；2 - 优选（默认）
	ShowMode int `json:"show_mode,omitempty"`
	// SiteType 预约广告
	// 1 - IOS预约；2 - ANDROID 预约；【仅在计划类型为收集营销线索，web_uri_type = 2 时生效】
	SiteType int `json:"site_type,omitempty"`
	// SmartCover 程序化创意 2.0 智能抽帧
	// 是否开启智能抽帧
	SmartCover *bool `json:"smart_cover,omitempty"`
	// AssetMining 程序化创意 2.0 素材挖掘
	// 是否开启历史素材挖掘
	AssetMining *bool `json:"asset_mining,omitempty"`
	// ConsultID 咨询组件
	// 1、仅可被用于线索类计划下的 unit；2、仅当落地页使用了建站落地页或者建站程序化落地页时可使用；3、注意本字段不可被更新；4、本属性不可与附加表单组件(component_id)同时使用；通过「/rest/openapi/v2/lp/consult/list」接口获取
	ConsultID uint64 `json:"consult_id,omitempty"`
	// AdvCardOption 高级创意开关
	// 0：关闭 1:开启
	AdvCardOption *int `json:"adv_card_option,omitempty"`
	// AdvCardList 绑定卡片
	// card_type=100 为 1 个；card_type=101 为 3 个；card_type=102 为 3 个；card_type=103 为 1 个；通过接口「/rest/openapi/v1/asset/adv_card/list」获取
	AdvCardList []uint64 `json:"adv_card_list,omitempty"`
	// PlayableID 试玩 ID
	// 可选字段，开启试玩时存在，否则不存在。当用户上传试玩素材成功时返回，用于之后对于广告组中试玩创意的编辑操作。通过接口「 /rest/openapi/gw/dsp/v1/playable/list 」 获取
	PlayableID uint64 `json:"playable_id,omitempty"`
	// PlayButton 试玩按钮文字内容
	// 开启试玩时存在，否则不存在，示例按钮内容如下：1：立即试玩；2：试玩一下；3：立即体验；4：免装试玩；5：免装体验。启用试玩时：默认“立即试玩”，未启用试玩时：内容为空。通过接口「/rest/openapi/gw/dsp/v1/playable/play_buttons」获取
	PlayButton string `json:"play_button,omitempty"`
	// DpaUnitParam DPA 相关商品信息
	// 当计划类型为商品库推广时必填。
	DpaUnitParam *DpaUnitParam `json:"dpa_unit_param,omitempty"`
	// JingleBellID 小铃铛组件id
	// 计划 campaignType=16 粉丝直播推广时必填写
	JingleBellID uint64 `json:"jingle_bell_id,omitempty"`
	// LiveUserID 主播id
	// 计划 campaignType=16 粉丝直播推广时必填写, 计划 campaignType=30 短剧推广时，必填，值为短剧作者ID
	LiveUserID uint64 `json:"live_user_id,omitempty"`
	// ConversionType 转化途径
	// 计划 campaignType=16 粉丝直播推广时必填写 6
	ConversionType int `json:"conversion_type,omitempty"`
	// ExtendSearch 智能扩词开启状态
	// false：关闭，true：开启，不填默认为false
	ExtendSearch *bool `json:"extend_search,omitempty"`
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
	QuickSearch *int `json:"quick_search,omitempty"`
	// TargetExplore 搜索人群探索
	// 0：不开启；1：开启；开启后系统将基于用户搜索行为探索定向人群以外的更多优质人群，仅在 quick_search = 1时生效；
	TargetExplore *int `json:"target_explore,omitempty"`
	// NegativeWordParam 搜索广告否词
	// 设置否词必填
	NegativeWordParam *NegativeWordParam `json:"negative_word_param,omitempty"`
	// PackageID 新版应用中心，应用包ID
	// 使用新版应用中心投放，该字段必填；老版本应用非必填 获取母包接口文档 获取分包接口文档
	PackageID uint64 `json:"package_id,omitempty"`
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

// Url implement PostRequest interface
func (r UpdateRequest) Url() string {
	return "gw/dsp/unit/update"
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateResponse 修改广告组 API Response
type UpdateResponse struct {
	// UnitID 广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
}
