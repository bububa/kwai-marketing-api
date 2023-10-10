package unit

import "encoding/json"

type Unit struct {
	CampaignID            uint64                    `json:"campaign_id"`             // 广告计划 ID
	UnitID                uint64                    `json:"unit_id"`                 // 广告组 ID
	UnitName              string                    `json:"unit_name"`               // 广告组名称
	PutStatus             int                       `json:"put_status"`              // 投放状态（操作结果） 1：投放中；2：暂停 3：删除
	Status                int                       `json:"status"`                  // 广告组状态（优先看看这个状态，计算结果） -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，，11：审核中，12：审核未通过，14：已结束，15：已暂停，17：组超预算，19：未达投放时间，20：有效，22：不在投放时段
	ReviewDetail          string                    `json:"review_detail"`           // 审核拒绝理由
	StudyStatus           int                       `json:"study_status"`            // 学习期 1:学习中,2:学习成功,3:学习失败
	CompensateStatus      int                       `json:"compensate_status"`       // 赔付状态 0=不需要赔付(不需要展示赔付标志)，1=成本保障生效中，2=成本保证确认中，3=已赔付完成，4=已失效
	BidType               int                       `json:"bid_type"`                // 出价类型 1：CPM，2：CPC，6：OCPC(使用 OCPC 代表 OCPX)，10：OCPM，20：eCPC
	Bid                   int64                     `json:"bid"`                     // 出价 单位：厘
	CPABid                int64                     `json:"cpa_bid"`                 // OCPC 出价 单位：厘
	OCPXActionType        int                       `json:"ocpx_action_type"`        // 优化目标 0：未知，2：点击转化链接，10：曝光，11：点击，31：下载完成，53：提交线索，109：电话卡激活，137：量房，180：激活，190: 付费，191：首日 ROI，348：有效线索，383: 授信，384: 完件 715：微信复制;739:7 日付费次数;774:7 日ROI;810:激活付费
	DeepConversionType    int                       `json:"deep_conversion_type"`    // 深度转化目标 3:付费，7:次日留存，10:完件, 11:授信；13:添加购物车；14:提交订单；15:购买；44：有效线索；92：付费 roi；181：激活后24H次日留存；0：无；
	DeepConversionBid     int64                     `json:"deep_conversion_bid"`     //深度转化目标出价
	EnhanceConversionType int                       `json:"enhance_conversion_type"` // 增强目标
	ROIRatio              float64                   `json:"roi_ratio"`               // 付费 ROI 系数，优化目标为「首日 ROI」时必填：ROI 系数取值范围 ( 0,100 ] 最多支持到三位小数（如：0.066）
	SceneID               []int                     `json:"scene_id"`                // 广告位，1：优选广告位；2：按场景选择广告位-信息流广告（旧广告位，包含上下滑大屏广告）6：上下滑大屏广告；7：信息流广告（不包含上下滑大屏广告）24：激励视频；11：快看点场景
	UnitType              int                       `json:"unit_type"`               // 创意制作方式，4: 自定义; 7：程序化创意 2.0
	BeginTime             string                    `json:"begin_time"`              // 投放开始时间，格式：yyyy-MM-dd
	EndTime               *string                   `json:"end_time,omitempty"`      // 投放结束时间，格式：yyyy-MM-dd, 排期不限为 null
	Schedule              interface{}               `json:"schedule,omitempty"`
	DayBudget             int64                     `json:"day_budget"`                  // 单日预算，单位：厘
	DayBudgetSchedule     []int64                   `json:"day_budget_schedule"`         // 分日预算，单位：厘，单日预算和分日预算同时存在时，以分日预算为准，优先级高于day_budget
	ConvertID             int                       `json:"convert_id"`                  // 转化目标
	URLType               int                       `json:"url_type"`                    // url 类型，当计划类型为 3 时（获取电商下单）时有返回。1：淘宝商品短链 2：淘宝商品 itemID
	WebURIType            int                       `json:"web_uri_type"`                // url 类型，当计划类型为 5（收集销售线索）&使用建站时有返回：需使用魔力建站 不传默认 1，2：落地页
	URL                   string                    `json:"url"`                         // 落地页链接，计划类型是 2（提升应用安装）：返回应用下载地址；计划类型是 3（获取电商下单）：根据 url_type 返回相应信息；计划类型是 4（推广品牌活动）：返回落地页 url；计划类型是 5（收集销售线索）：返回落地页 url；计划类型是 5（收集销售线索）：建站 id
	SiteID                int64                     `json:"site_id"`                     // 建站ID/下载中间页ID，当 web_uri_type = 2 时表示建站 ID，必须为数字，通过「/rest/openapi/v2/lp/page/list」获取。计划类型是2（提升应用安装）且关联应用为安卓时，表示安卓下载中间页ID，通过「/rest/openapi/v2/lp/page/list」获取 "view_comps = 7" 的建站ID。
	GroupID               int64                     `json:"group_id"`                    // 程序化落地页ID，web_uri_type = 3 时表示程序化落地页ID，必须为数字，通过「/rest/openapi/gw/magicsite/v1/group/list」获取；
	PageGroup             PageGroup                 `json:"page_group_detail"`           // 程序化落地页信息，广告组ID绑定的程序化落地页组信息
	SchemaURI             string                    `json:"schema_uri"`                  // 调起链接，提升应用活跃营销目标的调起链接
	SchemaID              string                    `json:"schema_id"`                   // 微信小程序外部调起链接，目前只有收集营销线索计划下的联盟广告位该字段才有效
	AppID                 int64                     `json:"app_id"`                      // 应用 ID
	AppDownloadType       int                       `json:"app_download_type"`           // 应用下载方式 0：直接下载 1：落地页下载（仅在提升应用安装计划下生效）
	UseAppMarket          int                       `json:"use_app_market,omitempty"`    // 优先从系统应用商店下载 0：未设置 1：优先从系统应用商店下载使用
	AppStore              []string                  `json:"app_store,omitempty"`         // 应用商店列表
	DiverseData           DiverseData               `json:"diverse_data"`                // 应用信息，详情见下方
	ShowMode              int                       `json:"show_mode"`                   // 创意展现方式 0：未知 1：轮播 2：优选
	SiteType              int                       `json:"site_type,omitempty"`         // 预约广告 1:IOS 预约 缺省为不传或传 0
	SmartCover            bool                      `json:"smart_cover"`                 // 程序化创意 2.0 智能抽帧 是否开启智能抽帧
	AssetMining           bool                      `json:"asset_mining"`                // 程序化创意 2.0 素材挖掘 是否开启历史素材挖掘
	ConsultID             int64                     `json:"consult_id"`                  // 是否使用了咨询组件 0=未使用 1=使用；注，咨询组件仅在收集销售线索计划(campaign_type=5)下可用，且使用了咨询组件后，可用的行动号召按钮限于接口返回内容
	AdvCardOption         int                       `json:"adv_card_option"`             // 高级创意开关 0：关闭 1:开启
	AdvCardList           []interface{}             `json:"adv_card_list,omitempty"`     // 高级创意列表
	PlayableID            int64                     `json:"playable_id"`                 // 试玩 ID 可选字段，开启试玩时存在，否则不存在。当用户上传试玩素材成功时返回，用于之后对于广告组中试玩创意的编辑操作。
	PlayButton            string                    `json:"play_button"`                 // 试玩按钮文字内容 开启试玩时存在，否则不存在，示例按钮内容如下：1：立即试玩；2：试玩一下；3：立即体验；4：免装试玩；5：免装体验。启用试玩时：默认“立即试玩”，未启用试玩时：内容为空。
	PlayableOrientation   int                       `json:"playable_orientation"`        // 试玩素材的横竖适配 默认值为-1；0:横竖均可；1:竖屏；2:横屏
	PlayableFileName      string                    `json:"playable_file_name"`          // 试玩广告的文件名
	DpaUnitSubType        int                       `json:"dpa_unit_sub_type"`           // 商品广告类型：1-DPA，2-SDPA，3-动态商品卡
	LibraryID             int                       `json:"library_id"`                  // 商品库ID
	OuterID               string                    `json:"outer_id"`                    // 商品第三方ID
	DpaOuterIDs           []string                  `json:"dpa_outer_ids"`               // DPA商品ID集合
	DpaCategoryIDs        []string                  `json:"dpa_category_ids"`            // DPA商品类目集合
	ProductName           string                    `json:"product_name"`                // 商品名称
	ProductPrice          float64                   `json:"product_price"`               // 商品价格，单位：元
	ProductImage          string                    `json:"product_image"`               // 商品主图
	JingleBellID          int64                     `json:"jingle_bell_id"`              // 小铃铛组件ID
	LiveUserID            int64                     `json:"live_user_id"`                // 主播ID
	ExtendSearch          bool                      `json:"extend_search"`               // 智能扩词开启状态，false：关闭，true：开启
	CustomMiniAppData     CustomMiniAppData         `json:"custom_mini_app_data"`        // 广小程序营销目标小程序信息 推广快手小程序时必填
	Target                Target                    `json:"target"`                      // 定向数据
	TemplateID            int64                     `json:"template_id"`                 // 定向模板ID
	OuterLoopNative       int                       `json:"outer_loop_native,omitempty"` // 是否开启原生
	QuickSearch           int                       `json:"quick_search,omitempty"`      // 是否开启快投
	TargetExplore         int                       `json:"target_explore,omitempty"`    // 是否开启搜索人群探索
	NegativeWord          AdMarketNegativeWordParam `json:"negative_word,omitempty"`     // 搜索广告否词
	BackflowForecast      json.RawMessage           `json:"backflow_forecast,omitempty"` // 回流预估数据
	AdType                int                       `json:"ad_type,omitempty"`           // 广告计划类型
	UnitSource            int                       `json:"unit_source,omitempty"`       // 广告组来源
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UdpateTime 最后修改时间
	UpdateTime string `json:"update_time,omitempty"`
	PackageID  int64  `json:"package_id,omitempty"` // 应用包ID
}

type PageReviewDetail struct {
	PageId       int64  `json:"page_id"`       // 页面 ID
	ReviewStatus string `json:"review_status"` // 落地页审核状态, 1-审核中 2-审核通过 3-审核拒绝
	Url          string `json:"url"`           // 落地页URL信息
	ReviewDetail string `json:"review_detail"` // 落地页的审核拒绝理由
}

type PageGroup struct {
	GroupId          int64              `json:"group_id"`           // 程序化落地页组 ID
	GroupName        string             `json:"group_name"`         // 程序化落地页组名称
	PageReviewDetail []PageReviewDetail `json:"page_review_detail"` // 程序化落地页组下的页面信息
}

type DiverseData struct {
	AppName        string `json:"app_name"`         // 应用名称
	AppPackageName string `json:"app_package_name"` // 应用包名
	DeviceOSType   int    `json:"device_os_type"`   // 应用操作系统类型，0：未知，1：ANDROID，2：IOS
}

type AdMarketNegativeWordParam struct {
	ExactWords  []string `json:"exact_words"`  // 精确否定词，搜索广告新增
	PhraseWords []string `json:"phrase_words"` // 短语否定词，搜索广告新增
}

type CustomMiniAppData struct {
	//mini_app_id_platform	string	选填	小程序APPID	计划 campaignType=19 推广快手小程序时必填，不能超过30个字符。 注：当账户加了小游戏推广白名单时，不需要填写。
	MiniAppIdPlatform string `json:"mini_app_id_platform,omitempty"`
	//bootstrap_page	string	选填	小程序启动页面	跳转小程序时将直接进入该页面，计划 campaignType=19 小程序推广时必填。例：page/app/index。注：当账户加了小游戏推广白名单时，必填。
	BootstrapPage string `json:"bootstrap_page,omitempty"`
	//bootstrap_params	string	选填	小程序启动参数	启动参数可以追踪页面来源，计划 campaignType=19 小程序推广时选填。例：name=lilei&age=18。注：当账户加了小游戏推广白名单时，不需要填写。
	BootstrapParams string `json:"bootstrap_params,omitempty"`
}
