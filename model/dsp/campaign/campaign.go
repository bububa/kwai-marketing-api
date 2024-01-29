package campaign

// Campaign 广告计划
type Campaign struct {
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// PutStatus 投放状态（操作结果）
	// 1：投放中；2：暂停 3：删除
	PutStatus int `json:"put_status,omitempty"`
	// Status 计划状态
	// 1：广告计划已暂停；3：广告计划超预算；4：有效；5：广告计划已删除；6：账户余额不足； -2：不限
	Status int `json:"status,omitempty"`
	// DayBudget 当日预算
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule 分日预算
	// 优先级高于day_budget
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
	// CampaignType 营销目标类型
	CampaignType int `json:"campaign_type"`
	// CampaignSubType 计划子类型
	// 商品库推广计划类型的老数据会返回【后续会进行下线】。4：DPA，5：SDPA
	CampaignSubType int `json:"campaign_sub_type"`
	// AdType 广告计划类型
	// 0:信息流，1:搜索
	AdType int `json:"ad_type,omitempty"`
	// BidType 出价类型
	BidType int `json:"bid_type,omitempty"`
	// AutoAdjust 自动调控开关
	// 0:关闭，1:开启
	AutoAdjust int `json:"auto_adjust,omitempty"`
	// AutoBuild 自动基建开关
	// 0:关闭，1:开启
	AutoBuild int `json:"auto_build,omitempty"`
	// AutoBuildNameRule 自动基建广告命名规则
	AutoBuildNameRule *AutoBuildNameRule `json:"auto_build_name_rule,omitempty"`
	// AutoManage 智能投放开关
	// 0:关闭，1:开启
	AutoManage int `json:"auto_manage,omitempty"`
	// CampaignOcpxActionType 智能模式下的优化目标
	CampaignOcpxActionType int `json:"campaign_ocpx_action_type,omitempty"`
	// CampaignOcpxActionTypeName 智能模式下的优化目标名称
	CampaignOcpxActionTypeName string `json:"campaign_ocpx_action_type_name,omitempty"`
	// CampaignDeepConversionType 智能模式下的深度优化目标
	CampaignDeepConversionType int `json:"campaign_deep_conversion_type,omitempty"`
	// CampaignDeepConversionTypeName 智能模式下的深度优化目标名称
	CampaignDeepConversionTypeName string `json:"campaign_deep_conversion_type_name,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	// CapRoiRatio cost cap的roi约束(广告成本约束-ROI约束)
	// 范围：0~100
	CapRoiRatio float64 `json:"cap_roi_ratio,omitempty"`
	// CapBid cost cap的广告成本约束(广告成本约束-非ROI单/双约束)
	// 单位：厘，范围：0~8000元
	CapBid int64 `json:"cap_bid,omitempty"`
	// ÇonstraintCpa 浅层成本约束(广告成本约束-非ROI双约束)
	// 单位：厘，范围：0~8000元
	ConstraintCpa int64 `json:"constraint_cpa,omitempty"`
}

// AutoBuildNameRule 自动基建命名规则
type AutoBuildNameRule struct {
	// UnitNameRule 广告组名称命名规则
	// 必须同时包含[日期]和[序号]宏变量，eg: 系统自动搭建_[日期][序号]
	UnitNameRule string `json:"unit_name_rule,omitempty"`
	// CreativeNameRule 广告创意名称命名规则
	// 必须同时包含[日期]和[序号]宏变量，eg: 系统自动搭建_[日期][序号]
	CreativeNameRule string `json:"creative_name_rule,omitempty"`
}
