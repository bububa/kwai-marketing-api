package campaign

import "github.com/bububa/kwai-marketing-api/model"

// CreateRequest 创建广告计划 API Request
type CreateRequest struct {
	// AdvertiserID 账号ID，在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignName 广告计划名称
	// 长度为 1-100 个字符，同一个账号下面计划名称不能重复
	CampaignName string `json:"campaign_name,omitempty"`
	// Type 营销目标类型
	// 2：提升应用安装；3：获取电商下单；4：推广品牌活动；5：收集销售线索；7：提高应用活跃；9：商品库推广（此营销目标下创建的默认为DPA广告）；16：粉丝/直播推广；19：小程序推广 30：快手号-短剧推广
	Type int `json:"type,omitempty"`
	// DailyBudget 单日预算金额
	// 单位：厘，指定 0 表示预算不限，默认为 0；不小于 500 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该计划当日花费的 120% 和修改前预算的较小者，与 day_budget_schedule 不能同时传，不能低于该计划下任一广告组出价。当 bid_type = 1时，day_budget 或者 budget_schedule 二选一必填
	DailyBudget int64 `json:"daily_budget,omitempty"`
	// DailyBudgetSchedule 分日预算
	// 单位：厘，指定 0 表示预算不限，默认为 0；每天不小于 500 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该计划当日花费的 120% 和修改前预算的较小者，与 day_budget 不能同时传，均不能低于该计划下任一广告组出价。事例：时间周期为周一到周日，样例："day_budget_schedule":[11110000,22220000,0,0,0,0,0]，优先级高于day_budget。当 bid_type = 1时，day_budget 或者 budget_schedule 二选一必填
	DailyBudgetSchedule int64 `json:"daily_budget_schedule,omitempty"`
	// AdType 0:信息流，1:搜索；不填默认信息流
	AdType int `json:"ad_type,omitempty"`
	// BidType 出价类型; 0:默认1:最大转化（添加后不可修改）
	BidType int `json:"bid_type,omitempty"`
	// AutoAdjust 自动调控开关
	// 0：关闭，1：开启
	AutoAdjust int `json:"auto_adjust,omitempty"`
	// AutoBuild 自动基建开关
	// 0：关闭，1：开启
	AutoBuild int `json:"auto_build,omitempty"`
	// AutoBuildNameRule 自动基建命名规则
	// 仅在auto_build为1时，此字段生效，开启自动基建时必填命名规则，组、创意命名规则均必须同时包含[日期]和[序号]宏变量【注：白名单功能】
	AutoBuildNameRule *AutoBuildNameRule `json:"auto_build_name_rule,omitempty"`
	// CapRoiRatio cost cap的roi约束(广告成本约束-ROI约束)
	// 默认0表示“不限”，新建计划时仅支持“不限”，当计划下存在广告组选择ROI出价，编辑时支持修为该约束
	CapRoiRatio float64 `json:"cap_roi_ratio,omitempty"`
	//  CapBid cost cap的成本约束(广告成本约束-非ROI单/双约束)
	// 默认0表示“不限”，新建计划时仅支持“不限”，当计划下存在广告组选择非ROI出价，编辑时支持修为该约束
	CapBid int64 `json:"cap_bid,omitempty"`
	// ConstraintCpa 浅层成本约束(广告成本约束-非ROI双约束)
	// 默认0表示“不限”，新建计划时仅支持“不限”，当计划下存在广告组选择自然日次日留存，编辑时支持修为该约束
	ConstraintCpa int64 `json:"constrait_cpa,omitempty"`
	// AutoManage 全自动投放开关 (UAX，支持UAA和UAL)
	// 0：关闭，1：开启 【注：此字段设置为开启时， 需要开启auto_adjust和auto_build 字段】 根据行业区分UAA&UAL具体详情可联系对应运营加白
	AutoManage int `json:"auto_manage,omitempty"`
}

// Url implement PostRequest interface
func (r CreateRequest) Url() string {
	return "gw/dsp/campaign/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CreateResponse 创建广告计划 API Response
type CreateResponse struct {
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
