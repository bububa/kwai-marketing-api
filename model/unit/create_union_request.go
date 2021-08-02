package unit

import (
	"encoding/json"

	"github.com/bububa/kwai-marketing-api/model/target"
)

// CreateUnionRequest 创建联盟定投广告组APIRequest
type CreateUnionRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignID int64 `json:"campaign_id,omitempty"`
	// UnitName 广告组名称; 长度为1-100个字符，同一个计划下的广告组名称不能重复
	UnitName string `json:"unit_name,omitempty"`
	// BidType 优化目标出价类型; 1: CPM 2: CPC 10:OCPM 4: CPA只有计划类型为3，url_type=4才支持，而且ocpx_action_type只能为表单提交53)
	BidType int `json:"bid_type,omitempty"`
	// Bid 出价; bid_type为CPC/eCPC/CPM时该字段必填，单位：厘，不得低于0.2元，不得高于100元，不得高于组预算
	Bid int64 `json:"bid,omitempty"`
	// CpaBid 出价; bid_type是OCPM时该字段必填，单位：厘，ocpx_action_type为2时，不得低于1元，不得高于3000元；ocpx_action_type为180时，不得低于1元，不得高于3000元，ocpx_action_type为53时，不得低于5元，不得高于3000元；不得高于组预算
	CpaBid int64 `json:"cpa_bid,omitempty"`
	// OcpxActionType 优化目标;
	// bid_type是OCPM时该字段必填；2：行为数（actionBar点击）；180：激活数；53：表单数；190: 付费；191：首日ROI；324：唤起应用（campaign_type=7+主站+加白）；348：有效线索；383: 授信；384: 完件394：订单提交；（计划类型获取电商下单，非金牛）；396：注册（联盟暂不支持）；715：微信复制优化目标；（campaign_type=5&bid_type=10）；716：多转化事件(campaign_type=5&bid_type=10&使用魔力建站&线索转化类组件>=2种)；717：广告观看次数(campaign_type=2&bid_type=10&appi与转化追踪工具联调过或者该广告主的在相应事件的回传数>=1）需通过接口「rest/openapi/v1/ad_unit/ocpc/conversion_infos」获取是否可以选择「激活」、「表单转化」、「付费」、「授信」、「完件」的优化目标;731：广告观看5次；732：广告观看10次；733：广告观看20次；773：关键行为；774：7日ROI；72：小店通商品和主页推广，转化目标“涨粉”；392：小店通商品推广，优化目标“转化数”，转化目标“商品访问”；395：小店通商品和主页推广，转化目标“订单支付”，62：直播观看数,（小店直播推广&商品推广支持，计划type=14、13）;192：直播推广ROI；小店直播推广&商品推广支持，计划type=14、13）;403：近似购买（小店直播推广&商品推广支持，计划type=14、13）。需通过接口「获取可选的深度转化类型」，接口is_order_paied字段返回值为1才可以使用;
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
	// DeepConversionType 深度转化目标;
	// 3:付费，7:次日留存，10:完件, 11:授信；13:添加购物车；14:提交订单；15:购买；44：有效线索；92：付费roi。可通过下方深度转化信息查询接口获取当前账户可选的深度转化目标
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// DeepConversionBid 深度转化目标出价; 单位：厘，仅当deep_conversion_type可用且账户满足白名单时选填，出价需大于cpa_bid，小于min{组预算，3000元}，不得高于组预算，不支持从0改为其他值，也不支持从其他值改为0
	DeepConversionBid int64 `json:"deep_conversion_bid,omitempty"`
	// SceneID 资源位置; 5：联盟广告；10：联盟优选 3、6、7、10可多选；当账户在联盟优选白名单中且campaign_type = 2/3/5/7，可传10；；当账户不在联盟优选白名单中或campaign_type不等于2/3/5/7时，不可传10；
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
	// ConvertID 转化目标ID; 计划类型为2时不支持；计划类型为5时仅支持【JS代码】的转化目标
	ConvertID int `json:"convert_id,omitempty"`
	// URL 投放链接; 当计划类型是3/4/5时必填；长度不超过1000字符；计划类型是3（获取电商下单）：金牛商品ID（必须为数字）；计划类型是4（推广品牌活动）：落地页URL；计划类型是5（收集销售线索）：落地页URL；计划类型是5（收集销售线索）：建站ID，通过/rest/openapi/v2/lp/page/list获取。「房地产」「家装建材」「招商加盟」三个二级行业【收集销售线索】目标下隐藏客户自有链接填写入口。
	URL string `json:"url,omitempty"`
	// AppID 应用ID; 当计划类型为2时必填，可通过应用列表接口获取应用ID；为9时且detail_unit_type为0、2时必填
	AppID int64 `json:"app_id,omitempty"`
	// ShowMode 创意展现方式; 1 - 轮播；2 - 优选
	ShowMode int `json:"show_mode,omitempty"`
	// Speed 投放方式; 1 - 加速投放；2 - 平滑投放；3-优先低成本（投放时间范围只可为全天；预算不可为不限或空）
	Speed int `json:"speed,omitempty"`
	// PlayableOrientation 试玩素材的横竖适配; 默认值为-1；0:横竖均可；1:竖屏；2:横屏
	PlayableOrientation int `json:"playable_orientation,omitempty"`
	// PlayableUrl 试玩的url
	PlayableUrl string `json:"playable_url,omitempty"`
	// PlayableSwitch 试玩广告的开关; 默认值为0；1:关闭；2:开启
	PlayableSwitch int `json:"playalbe_switch_omitempty"`
	// Target 定向数据
	Target *target.Target `json:"target,omitempty"`
}

// Url implement PostRequest interface
func (r CreateUnionRequest) Url() string {
	return "v2/ad_unit/create"
}

// Encode implement PostRequest interface
func (r CreateUnionRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
