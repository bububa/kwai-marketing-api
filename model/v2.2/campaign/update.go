package campaign

import "github.com/bububa/kwai-marketing-api/model"

// UpdateRequest 更新广告计划 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignId uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告计划名称; 长度为1-100个字符，计划名称不能重复
	CampaignName string `json:"campaign_name,omitempty"`
	// Type 计划类型; 2：提升应用安装；3：获取电商下单；4：推广品牌活动；5：收集销售线索；7：提高应用活跃；9：商品库推广（此营销目标下创建的默认为DPA广告）；16：粉丝/直播推广；19：小程序推广
	Type int `json:"type,omitempty"`
	// DayBudget 单日预算金额; 单位：厘，指定0表示预算不限，默认为0；不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget_schedule不能同时传，不能低于该计划下任一广告组出价。当bid_type=1时，day_budget或者budget_schedule二选一必填
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算; 单位：厘，指定0表示预算不限，默认为0；每天不小于500元，不超过100000000元，仅支持输入数字；修改预算不得低于该计划当日花费的120%，与day_budget不能同时传，均不能低于该计划下任一广告组出价。事例：时间周期为周一到周日，样例："day_budget_schedule":[11110000,22220000,0,0,0,0,0]，优先级高于day_budget。当bid_type=1时，day_budget或者budget_schedule二选一必填
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
	// AdType 广告计划类型; 0:信息流，1:搜索；不填默认信息流
	AdType int `json:"ad_type,omitempty"`
	// BidType 出价类型; 0:默认1:最大转化（添加后不可修改）
	BidType int `json:"bid_type,omitempty"`
	// AutoAdjust 自动调控开关; 0：关闭，1：开启【注：此字段设置为关闭时，auto_build字段也必须为关闭，白名单功能】
	AutoAdjust int `json:"auto_adjust,omitempty"`
	// AutoBuild 自动基建开关; 0：关闭，1：开启【注：此字段设置为开启时，auto_adjust字段也必须为开启，白名单功能】
	AutoBuild int `json:"auto_build,omitempty"`
	// AutoBuildNameRule 自动基建命名规则; 仅在auto_build为1时，此字段生效，开启自动基建时必填命名规则，组、创意命名规则均必须同时包含[日期]和[序号]宏变量【注：白名单功能】
	AutoBuildNameRule *AutoBuildNameRule `json:"auto_build_name_rule,omitempty"`
	// auto_manage	int	选填	是	智能投放开关	0：关闭，1：开启。【注：此字段修改为关闭时，需同时关闭 auto_adjust、auto_build 字段。此字段修改为开启时，该计划下必须存在在投广告组且campaign_ocpx_action_type、campaign_deep_conversion_type必须设置为与任一在投广告组优化目标和深层优化目标值一致，白名单功能】
	AutoManage int `json:"auto_manage,omitempty"`
	// campaign_ocpx_action_type	int	选填	是	智能投放优化目标	auto_manage 设置为开启时可填，必须与该计划下任一广告组的优化目标保持一致
	CampaignOcpxActionType int `json:"campaign_ocpx_action_type,omitempty"`
	// campaign_deep_conversion_type	int	选填	是	智能投放深度优化目标	auto_manage 设置为开启时可填，必须与该计划下任一广告组的深度优化目标保持一致
	CampaignDeepConversionType int `json:"campaign_deep_conversion_type,omitempty"`
	// PutStatus 1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateRequest) Url() string {
	return "gw/dsp/campaign/update"
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateResponse 更新广告计划 API Response
type UpdateResponse struct {
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
