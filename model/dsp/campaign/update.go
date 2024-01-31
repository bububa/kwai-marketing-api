package campaign

import "github.com/bububa/kwai-marketing-api/model"

// UpdateRequest 修改广告计划 API Request
type UpdateRequest struct {
	// AdvertiserID 账号ID，在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告计划名称
	// 长度为 1-100 个字符，同一个账号下面计划名称不能重复
	CampaignName string `json:"campaign_name,omitempty"`
	// DayBudget 单日预算金额
	// 单位：厘，指定 0 表示预算不限，默认为 0；不小于 500 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该计划当日花费的 120% 和修改前预算的较小者，与 day_budget_schedule 不能同时传，不能低于该计划下任一广告组出价。当 bid_type = 1时，day_budget 或者 budget_schedule 二选一必填
	DayBudget int64 `json:"day_budget,omitempty"`
	// DailyBudgetSchedule 分日预算
	// 单位：厘，指定 0 表示预算不限，默认为 0；每天不小于 500 元，不超过 100000000 元，仅支持输入数字；修改预算不得低于该计划当日花费的 120% 和修改前预算的较小者，与 day_budget 不能同时传，均不能低于该计划下任一广告组出价。事例：时间周期为周一到周日，样例："day_budget_schedule":[11110000,22220000,0,0,0,0,0]，优先级高于day_budget。当 bid_type = 1时，day_budget 或者 budget_schedule 二选一必填
	DayBudgetSchedule int64 `json:"day_budget_schedule,omitempty"`
	// AutoAdjust 自动调控开关
	// 0：关闭，1：开启
	AutoAdjust *int `json:"auto_adjust,omitempty"`
	// AutoBuild 自动基建开关
	// 0：关闭，1：开启
	AutoBuild *int `json:"auto_build,omitempty"`
	// AutoBuildNameRule 自动基建命名规则
	// 仅在auto_build为1时，此字段生效，开启自动基建时必填命名规则，组、创意命名规则均必须同时包含[日期]和[序号]宏变量【注：白名单功能】
	AutoBuildNameRule *AutoBuildNameRule `json:"auto_build_name_rule,omitempty"`
	// CapRoiRatio cost cap的roi约束(广告成本约束-ROI约束)
	// 默认0表示“不限”，新建计划时仅支持“不限”，当计划下存在广告组选择ROI出价，编辑时支持修为该约束
	CapRoiRatio *float64 `json:"cap_roi_ratio,omitempty"`
	//  CapBid cost cap的成本约束(广告成本约束-非ROI单/双约束)
	// 默认0表示“不限”，新建计划时仅支持“不限”，当计划下存在广告组选择非ROI出价，编辑时支持修为该约束
	CapBid *int64 `json:"cap_bid,omitempty"`
	// ConstraintCpa 浅层成本约束(广告成本约束-非ROI双约束)
	// 默认0表示“不限”，新建计划时仅支持“不限”，当计划下存在广告组选择自然日次日留存，编辑时支持修为该约束
	ConstraintCpa *int64 `json:"constrait_cpa,omitempty"`
	// PutStatus 投放状态
	// 1-投放中、2-暂停、3-删除
	PutStatus int `json:"put_status,omitempty"`
	// AutoManage 全自动投放开关 (UAX，支持UAA和UAL)
	// 0：关闭，1：开启 【注：此字段设置为开启时， 需要开启auto_adjust和auto_build 字段】 根据行业区分UAA&UAL具体详情可联系对应运营加白
	AutoManage *int `json:"auto_manage,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateRequest) Url() string {
	return "gw/dsp/campaign/update"
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateResponse 修改广告计划 API Response
type UpdateResponse struct {
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
