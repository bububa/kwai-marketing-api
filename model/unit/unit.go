package unit

import "github.com/Shinku-Chen/kwai-marketing-api/model/target"

// Unit 广告组
type Unit struct {
	// CampaignID 广告计划ID
	CampaignID int64 `json:"campaign_id,omitempty"`
	// UnitID 广告组ID
	Unit int64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// Status 广告组状态（优先看看这个状态，计算结果）; -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，，11：审核中，12：审核未通过，14：已结束，15：已暂停，17：组超预算，19：未达投放时间，20：有效
	Status int `json:"status,omitempty"`
	// PutStatus 投放状态（操作结果）;1：投放中；2：暂停 3：删除
	PutStatus int `json:"put_status,omitempty"`
	// CreateChannel 创建渠道; 0：投放后台创建；1：Marketing API创建
	CreateChannel int `json:"create_channel,omitempty"`
	// ReviewDetail 审核拒绝理由;
	ReviewDetail string `json:"review_detail,omitempty"`
	// StudyStatus 学习期;1:学习中,2:学习成功,3:学习失败
	StudyStatus int `json:"study_status,omitempty"`
	// CompensateStatus 赔付状态; 0=不需要赔付(不需要展示赔付标志)，1=成本保障生效中，2=成本保证确认中，3=已赔付完成，4=已失效
	CompensateStatus int `json:"compensate_status,omitempty"`
	// BidType 出价类型; 1：CPM，2：CPC，6：OCPC(使用OCPC代表OCPX)，10：OCPM，20：eCPC
	BidType int `json:"bid_type,omitempty"`
	// Bid 出价; 单位：厘
	Bid int64 `json:"bid,omitempty"`
	// CpaBid OCPC出价; 单位：厘
	CpaBid int64 `json:"cpa_bid,omitempty"`
	// SmartBid 优先低成本是否自动出价; 0：手动出价，1：自动出价
	SmartBid int `json:"smart_bid,omitempty"`
	// OcpxActionType 优化目标; 0：未知，2：点击转化链接，10：曝光，11：点击，31：下载完成，53：提交线索，180：激活，190: 付费，191：首日ROI，348：有效线索，383: 授信，384: 完件 715：微信复制
	OcpxActionType int `json:"ocpx_action_type"`
	// DeepConversionType 深度转化目标; 3: 付费，7: 次日留存，10: 完件, 11: 授信 ，0：无
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// DeepConversionBid 深度转化目标出价; 白名单功能，单位：厘
	DeepConversionBid int64 `json:"deep_conversion_bid,omitempty"`
	// DayBudget 单日预算; 单位：厘
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，单日预算和分日预算同时存在时，以分日预算为准
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
	// RoiRatio 付费ROI系数; 优化目标为「首日ROI」时必填：ROI系数取值范围 ( 0,100 ] 最多支持到三位小数（如：0.066）
	RoiRatio float64 `json:"roi_ratio,omitempty"`
	// Speed 投放方式; 0：未知，1：正常投放，2：平滑投放，3：优先低成本。
	Speed int `json:"speed,omitempty"`
	// BeginTime 投放开始时间; 格式：yyyy-MM-dd
	BeginTime string `json:"begin_time,omitempty"`
	// EndTime 投放结束时间; 格式：yyyy-MM-dd,排期不限为null
	EndTime string `json:"end_time,omitempty"`
	// Schedule 投放时段
	Schedule *Schedule `json:"schedule,omitempty"` // 不投放的时段为null，详请见下方，历史字段，即将废弃
	// ScheduleTime 投放时段; 24*7的字符串，0为不投放，1为投放，例如：0010000000001....0000
	ScheduleTime string `json:"schedule_time,omitempty"`
	// SceneID 广告位; 1：优选广告位；2：按场景选择广告位-信息流广告（旧广告位，包含上下滑大屏广告）；3：视频播放页广告-便利贴广告；6：上下滑大屏广告；7：信息流广告（不包含上下滑大屏广告）10：激励视频；11：快看点场景
	SceneID []int `json:"scene_id,omitempty"`
	// ShowMode 创意展现方式; 0：未知，1：轮播，2：优选
	ShowMode int `json:"show_mode,omitempty"`
	// UnitType 创意制作方式; 4: 自定义；5：程序化创意 7：程序化创意2.0
	UnitType int `json:"unit_type,omitempty"`
	// ConvertID 转化目标
	ConvertID int `json:"convert_id,omitempty"`
	// UseAppMarket 优先从系统应用商店下载; 0：未设置 1：优先从系统应用商店下载使用
	UseAppMarket int `json:"use_app_market,omitempty"`
	// AppStore 应用商店列表
	AppStore []string `json:"app_store,omitempty"`
	// UrlType url类型; 当计划类型为3时（获取电商下单）时有返回。1：淘宝商品短链 2：淘宝商品itemID
	UrlType int `json:"url_type,omitempty"`
	// WebUriType url类型; 当计划类型为5（收集销售线索）&使用建站时有返回：需使用魔力建站 不传默认1，2：落地页
	WebUriType int `json:"web_uri_type,omitempty"`
	// Url 落地页链接; 计划类型是 2（提升应用安装）：返回应用下载地址；计划类型是 3（获取电商下单）：根据url_type返回相应信息；计划类型是 4（推广品牌活动）：返回落地页url计划类型是 5（收集销售线索）：返回落地页url计划类型是5（收集销售线索）：建站idl
	Url string `json:"url,omitempty"`
	// SchemaUrl 调起链接; 提升应用活跃营销目标的调起链接
	SchemaUrl string `json:"schema_url,omitempty"`
	// AppID 应用ID
	AppID int64 `json:"app_id,omitempty"`
	// AppIconUrl APP图标存储地址
	AppIconUrl string `json:"app_icon_url,omitempty"`
	// DiverseData 应用信息
	DiverseData *DiverseData `json:"diverse_data,omitempty"`
	// Target 定向数据
	Target *target.Target `json:"target,omitempty"`
	// SiteType 预约广告; 1:IOS预约 缺省为不传或传0
	SiteType int `json:"site_type,omitempty"`
	// GiftData 游戏礼包码; 联盟广告暂不支持
	GiftData *GiftData `json:"gift_data,omitempty"`
	// VideoLandingPage 是否使用落地页前置功能; true: 使用 false：不使用
	VideoLandingPage bool `json:"video_landing_page,omitempty"`
	// AutoTarget 智能定向; 若开启智能定向，定向部分仅保留地域+年龄+性别+排除人群+系统纬度的定向，其他定向纬度暂不支持（报错处理）true表示开启，false表示未开启
	AutoTarget bool `json:"auto_target,omitempty"`
	// SmartCover 程序化创意2.0智能抽帧; 是否开启智能抽帧
	SmartCover bool `json:"smart_cover,omitempty"`
	// AssetMining 程序化创意2.0素材挖掘; 是否开启历史素材挖掘
	AssetMining bool `json:"asset_mining,omitempty"`
	// AutoCreatePhoto 是否开启自动生成视频
	AutoCreatePhoto bool `json:"auto_create_photo,omitempty"`
	// UpdateTime 最后修改时间
	UpdateTime string `json:"update_time,omitempty"`
	// ItemID 电商关联Id (小店通); 1. merchantItemType为0时表示小店商品id; 2. merchantItemType为2时表示快手id
	ItemID int64 `json:"item_id,omitempty"`
	// FictionID 小说ID; 仅支持“提升应用安装”、“收集销售线索”以及“提高应用活跃”三种计划类型，且一旦绑定，不可修改。此参数仅是绑定小说，并非自动关联小说生成的落地页，如需推广小说生成的落地页，请使用小说ID获取其生成的建站落地页后将落地页ID一并传入即可（落地页ID传参与之前建站落地页ID字段一致）
	FictionID int64 `json:"fiction_id,omitempty"`
	// MerchantItemPutType 电商广告投放类型（小店通）; 0: 商品; 2: 个人主页
	MerchantItemPutType int `json:"merchant_item_put_type,omitempty"`
	// ConsultID 是否使用了咨询组件; 0=未使用，1=使用；注，咨询组件仅在收集销售线索计划(campaign_type=5)下可用，且使用了咨询组件后，可用的行动号召按钮限于接口返回内容
	ConsultID int `json:"consult_id,omitempty"`
	// AdvCardOption 高级创意开关; 0：关闭 1:开启
	AdvCardOption int `json:"adv_card_option,omitempty"`
	// AdvCardList 高级创意列表
	AdvCardList []AdvCard `json:"adv_card_list,omitempty"`
	// BackflowForecast backflow_cv_lower：回流预估值的下限；backflow_cv_upper：回流预估值的上限；backflow_timestamp：本次回流预估数据的时间戳，13位毫秒时间戳
	BackflowForecast *BackflowForecast `json:"backflow_forecast,omitempty"`
	// MerchandiseID 商品ID，且一旦绑定，不可修改; 此参数用于绑定商品（绑定商品类型受merchandise_type字段控制），与 fiction_id 字段互斥。merchandise_type=2，merchandise_id 为课程ID，仅支持“收集销售线索”计划类型，且一旦绑定不可修改
	MerchandiseID int64 `json:"merchandise_id,omitempty"`
	// MerchandiseType 课程类型; 与 merchandise_id 共同使用，merchandise_type=2，merchandise_id 为课程ID，仅支持“收集销售线索”计划类型，且一旦绑定不可修改
	MerchandiseType int `json:"merchandise_type,omitempty"`
	// PlayableOrientation 试玩素材的横竖适配; 默认值为-1；0:横竖均可；1:竖屏；2:横屏
	PlayableOrientation int `json:"playable_orientation,omitempty"`
	// PlayableUrl 试玩的url
	PlayableUrl string `json:"playable_url,omitempty"`
	// PlayableFileName 试玩广告的文件名
	PlayableFileName string `json:"playable_file_name,omitempty"`
	// PlayableSwitch 试玩广告的开关; 默认值为0；1:关闭；2:开启
	PlayableSwitch int `json:"playable_switch,omitempty"`
	// LibraryID 商品库ID; sdpa类型广告组才会存在
	LibraryID int64 `json:"library_id,omitempty"`
	// OuterID 商品ID; sdpa类型广告组才会存在
	OuterID string `json:"outer_id,omitempty"`
	// ProductName 商品名称; sdpa类型广告组才会存在
	ProductName string `json:"product_name,omitempty"`
	// ProductPrice 商品价格; 单位：元，sdpa类型广告组才会存在
	ProductPrice float64 `json:"product_price,omitempty"`
	// ProductImage 商品主图; sdpa类型广告组才会存在
	ProductImage string `json:"product_image,omitempty"`
	// IntentionTarget 行为意向-系统优选; 行为意向是否开启系统优选，智能定向和行为意向系统优选不能同时开启
	IntentionTarget bool `json:"intention_target,omitempty"`

	CreateTime  string `json:"create_time,omitempty"`
	TemplateId  int    `json:"template_id,omitempty"`
	SchemaUri   string `json:"schema_uri,omitempty"`
	ComponentId int    `json:"component_id,omitempty"`
	//SupportUnitIds      interface{} `json:"support_unit_ids"`
	UseSka          bool   `json:"use_ska,omitempty"`
	PlayableId      int64  `json:"playable_id,omitempty"`
	PlayButton      string `json:"play_button,omitempty"`
	ProductId       int64  `json:"product_id,omitempty"`
	SplashAdSwitch  bool   `json:"splash_ad_switch,omitempty"`
	PageGroupDetail string `json:"page_group_detail,omitempty"`
	AdType          int    `json:"ad_type,omitempty"`
	//ExtendSearch        interface{} `json:"extend_search"`
}
