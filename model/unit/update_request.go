package unit

import (
	"encoding/json"

	"github.com/bububa/kwai-marketing-api/model/target"
)

// UpdateRequest 修改广告组APIRequest
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitName 广告组名称; 长度为1-100个字符，同一个计划下的广告组名称不能重复
	UnitName string `json:"unit_name,omitempty"`
	// PutStatus 广告组的投放状态; 不填时，创建的广告组为投放状态；填2时，创建的广告组为暂停状态，其他值非法
	PutStatus int `json:"put_status,omitempty"`
	// TemplateID 定向模板id
	TemplateID uint64 `json:"template_id,omitempty"`
	// BidType 优化目标出价类型; 1: CPM 2: CPC 10:OCPM 4: CPA只有计划类型为3，url_type=4才支持，而且ocpx_action_type只能为表单提交53)
	BidType int `json:"bid_type"`
	// UseAppMarket 优先从系统应用商店下载; 下载模式：允许值：0，1；1：优先从系统应用商店下载使用，默认0。仅提升应用安装计划类型且选择的app为安卓才支持。
	UseAppMarket int `json:"use_app_market,omitempty"`
	// AppStore 应用商店列表; 华为：huawei, OPPO：oppo VIVO：vivo 小米：xiaomi；魅族：meizu；锤子：smartisan 只支持计划类型为 2 且应用是 Android 类型的推广
	AppStore []string `json:"app_store,omitempty"`
	// Bid 出价; bid_type为CPC/eCPC/CPM时该字段必填，单位：厘，不得低于0.2元，不得高于100元，不得高于组预算
	Bid int64 `json:"bid,omitempty"`
	// CpaBid 出价; bid_type是OCPM时该字段必填，单位：厘，ocpx_action_type为2时，不得低于1元，不得高于3000元；ocpx_action_type为180时，不得低于1元，不得高于3000元，ocpx_action_type为53时，不得低于5元，不得高于3000元；不得高于组预算
	CpaBid int64 `json:"cpa_bid,omitempty"`
	// SmartBid 优先低成本出价(设置条件); 当 speed 等于 3 时可用 ；0：手动出价，1：自动出价；1、当 smart_bid 为 1 时，cpa_bid 或 deep_conversion_bid 都不能设置值；2、快享和联盟不支持优先低成本，直接报错；3、优先低成本的 bid_type 只支持OCPM；4、优先低成本 必须设置统一预算或分日预算
	SmartBid int `json:"smart_bid,omitempty"`
	// OcpxActionType 优化目标;
	// bid_type是OCPM时该字段必填；2：行为数（actionBar点击）；180：激活数；53：表单数；190: 付费；191：首日ROI；324：唤起应用（campaign_type=7+主站+加白）；348：有效线索；383: 授信；384: 完件394：订单提交；（计划类型获取电商下单，非金牛）；396：注册（联盟暂不支持）；715：微信复制优化目标；（campaign_type=5&bid_type=10）；716：多转化事件(campaign_type=5&bid_type=10&使用魔力建站&线索转化类组件>=2种)；717：广告观看次数(campaign_type=2&bid_type=10&appi与转化追踪工具联调过或者该广告主的在相应事件的回传数>=1）需通过接口「rest/openapi/v1/ad_unit/ocpc/conversion_infos」获取是否可以选择「激活」、「表单转化」、「付费」、「授信」、「完件」的优化目标;731：广告观看5次；732：广告观看10次；733：广告观看20次；773：关键行为；774：7日ROI；72：小店通商品和主页推广，转化目标“涨粉”；392：小店通商品推广，优化目标“转化数”，转化目标“商品访问”；395：小店通商品和主页推广，转化目标“订单支付”，62：直播观看数,（小店直播推广&商品推广支持，计划type=14、13）;192：直播推广ROI；小店直播推广&商品推广支持，计划type=14、13）;403：近似购买（小店直播推广&商品推广支持，计划type=14、13）。需通过接口「获取可选的深度转化类型」，接口is_order_paied字段返回值为1才可以使用;
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
	// DeepConversionType 深度转化目标;
	// 3:付费，7:次日留存，10:完件, 11:授信；13:添加购物车；14:提交订单；15:购买；44：有效线索；92：付费roi。可通过下方深度转化信息查询接口获取当前账户可选的深度转化目标
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// RoiRatio 付费ROI系数; 优化目标为「首日ROI」时必填：ROI系数取值范围 ( 0,100 ] 最多支持到三位小数（如：0.066）
	RoiRatio float64 `json:"roi_ratio,omitempty"`
	// DeepConversionBid 深度转化目标出价; 单位：厘，仅当deep_conversion_type可用且账户满足白名单时选填，出价需大于cpa_bid，小于min{组预算，3000元}，不得高于组预算，不支持从0改为其他值，也不支持从其他值改为0
	DeepConversionBid int64 `json:"deep_conversion_bid,omitempty"`
	// SceneID 资源位置; 1：优选广告位（当账户在联盟优选白名单中且campaign_type=2/3/5/7时为主站&联盟优选广告位 ）；3：视频播放页广告-便利贴广告（不支持深度转化目标的优化）；5：联盟广告；6：上下滑大屏广告；7：信息流广告；10：联盟场景 3、6、7、10可多选;当账户在联盟优选白名单中且campaign_type = 2/3/5/7，可传10；;当账户不在联盟优选白名单中或campaign_type不等于2/3/5/7时，不可传10；
	SceneID []int `json:"scene_id,omitempty"`
	// UnitType 创意制作方式; 4: 自定义；5：程序化创意 7：程序化创意2.0
	UnitType int `json:"unit_type,omitempty"`
	// BeginTime 投放开始时间; 格式为yyyy-MM-dd，需大于等于当前时间
	BeginTime string `json:"begin_time,omitempty"`
	// EndTime 投放结束时间; 格式为yyyy-MM-dd，不传值表示从今天开始长期投放，传值表示设置开始时间和结束时间，且投放结束时间需大于等于投放开始时间
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时间段; 格式为24*7位字符串，且都为0和1，以一个小时为最小粒度，从周一零点开始至周日24点结束；0为不投放，1为投放，不传/全传1/全传0视为全时段投放
	ScheduleTime string `json:"schedule_time,omitempty"`
	// DayBudget 单日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于100元，不超过100000000元，仅支持输入数字；修改预算不得低于该广告组当日花费的120%，与day_budget_schedule不能同时传，均不能低于组出价
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget不能同时传，均不能低于该计划下任一广告组出价
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
	// ConvertID 转化目标ID; 可通过下方【转化目标查询接口】获得，不同计划类型需要对应各自的转化目标类型：提升应用安装(campaign_type=2) - 安卓：convert_type:3、7 / IOS：convert_type:7；推广品牌活动(campaign_type=4) / 收集销售线索(campaign_type=5)：convert_type:1、2
	ConvertID int `json:"convert_id,omitempty"`
	// UrlType url类型; 当计划类型为：3（获取电商下单）时必填：1 - 淘宝商品短链；2 - 淘宝商品itemID；4 - 金牛电商
	UrlType int `json:"url_type,omitempty"`
	// WebUrlType url类型; 当计划类型为5（收集销售线索）&使用建站时必填：需使用魔力建站；不传默认1，2：落地页
	WebUrlType int `json:"web_url_type,omitempty"`
	// URL 投放链接; 当计划类型是3/4/5时必填；长度不超过1000字符；计划类型是3（获取电商下单）：金牛商品ID（必须为数字）；计划类型是4（推广品牌活动）：落地页URL；计划类型是5（收集销售线索）：落地页URL；计划类型是5（收集销售线索）：建站ID，通过/rest/openapi/v2/lp/page/list获取。「房地产」「家装建材」「招商加盟」三个二级行业【收集销售线索】目标下隐藏客户自有链接填写入口。
	URL string `json:"url,omitempty"`
	// SchemaUrl 调起链接; 提升应用活跃营销目标的调起链接；应用推广下在已经安装app的用户手机上可以拉起app（需要运营加白）
	SchemaUrl string `json:"schema_url,omitempty"`
	// AppID 应用ID; 当计划类型为2时必填，可通过应用列表接口获取应用ID；为9时且detail_unit_type为0、2时必填
	AppID uint64 `json:"app_id,omitempty"`
	// ShowMode 创意展现方式; 1 - 轮播；2 - 优选
	ShowMode int `json:"show_mode,omitempty"`
	// Speed 投放方式; 1 - 加速投放；2 - 平滑投放；3-优先低成本（投放时间范围只可为全天；预算不可为不限或空）
	Speed int `json:"speed,omitempty"`
	// SiteType 预约广告; 1:IOS预约缺省为不传或传0
	SiteType int `json:"site_type,omitempty"`
	// GiftData 游戏礼包码; "gift_data": {}，仅支持计划类型为 2
	GiftData *GiftData `json:"gift_data,omitempty"`
	// VideoLandingPage 是否使用落地页前置功能; true: 使用 false：不使用，不填使用系统默认值（只支持双feed流，推广品牌活动-落地页url填写、获取电商下单-淘客短链url填写、获取电商下单-金牛活动页、获取销售线、获取销售线索）
	VideoLandingPage *bool `json:"video_landing_page,omitempty"`
	// AutoTarget 智能定向; 若开启智能定向，定向部分仅保留地域+年龄+性别+排除人群+系统纬度的定向，其他定向纬度暂不支持（报错处理）true表示开启，false表示未开启
	AutoTarget *bool `json:"auto_target,omitempty"`
	// AutoCreatePhoto 是否开启自动生成视频; 加白使用
	AutoCreatePhoto *bool `json:"auto_create_photo,omitempty"`
	// ItemID 电商关联Id (小店通); 1. merchantItemType为0时填写小店商品id 2. merchantItemType为2时不用填写，系统补充
	ItemID uint64 `json:"item_id,omitempty"`
	// MerchantItemPutType 电商广告投放类型（小店通）; 0: 商品 2: 个人主页
	MerchantItemPutType *int `json:"merchant_item_put_type,omitempty"`
	// FictionID 小说ID; 仅支持“提升应用安装”、“收集销售线索”以及“提高应用活跃”三种计划类型，且一旦绑定，不可修改。此参数仅是绑定小说，并非自动关联小说生成的落地页，如需推广小说生成的落地页，请使用小说ID获取其生成的建站落地页后将落地页ID一并传入即可（落地页ID传参与之前建站落地页ID字段一致）
	FictionID uint64 `json:"fiction_id,omitempty"`
	// SmartCover 程序化创意2.0智能抽帧; 是否开启智能抽帧; unit_type=7选填
	SmartCover *bool `json:"smart_cover,omitempty"`
	// AssetMining 程序化创意2.0素材挖掘; 是否开启历史素材挖掘
	AssetMining *bool `json:"asset_mining,omitempty"`
	// ConsultID 咨询组件id; 1、仅可被用于线索类计划下的unit；2、仅当落地页使用了建站落地页时可使用；3、注意本字段不可被更新；4、本属性不可与附加表单组件(component_id)同时使用
	ConsultID uint64 `json:"consult_id,omitempty"`
	// AdvCardOption 高级创意开关; 0：关闭 1:开启
	AdvCardOption *int `json:"adv_card_option,omitempty"`
	// AdvCardList 绑定卡片id; card_type=100为1个；card_type=101为3个；card_type=102为3个；card_type=103为1个
	AdvCardList []uint64 `json:"adv_card_list,omitempty"`
	// CardType 卡片类型; 100:图片卡片 101:多利益卡-图文 102：多利益卡-多标签 103：电商促销样式 104：快捷评论卡
	CardType int `json:"card_type,omitempty"`
	// MerchandiseID 商品ID，且一旦绑定，不可修改; 此参数用于绑定商品（绑定商品类型受merchandise_type字段控制），与 fiction_id 字段互斥。merchandise_type=2，merchandise_id 为课程ID，仅支持“收集销售线索”计划类型，且一旦绑定不可修改
	MerchandiseID uint64 `json:"merchandise_id,omitempty"`
	// MerchandiseType 选填; 与 merchandise_id 共同使用，merchandise_type=2，merchandise_id 为课程ID，仅支持“收集销售线索”计划类型，且一旦绑定不可修改
	MerchandiseType int `json:"merchandise_type,omitempty"`
	// IntentionTarget 行为意向-系统优选; 行为意向是否开启系统优选，智能定向和行为意向系统优选不能同时开启
	IntentionTarget *bool `json:"intention_target,omitempty"`
	// Target 定向数据
	Target *target.Target `json:"target,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateRequest) Url() string {
	return "v2/ad_unit/update"
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
