package campaign

// Campaign 包含广告计划信息
type Campaign struct {
	CampaignID                     int64             `json:"campaign_id"`                                  // 计划id
	CampaignName                   string            `json:"campaign_name,omitempty"`                      // 计划名称
	PutStatus                      int               `json:"put_status,omitempty"`                         // 投放状态（操作结果） 1：投放中；2：暂停；3：删除
	Status                         int               `json:"status,omitempty"`                             // 计划状态 1：广告计划已暂停；3：广告计划超预算；4：有效；5：广告计划已删除；6：账户余额不足；-2：不限
	DayBudget                      int64             `json:"day_budget,omitempty"`                         // 当日预算
	DayBudgetSchedule              []int64           `json:"day_budget_schedule,omitempty"`                // 分日预算 优先级高于day_budget
	CampaignType                   int               `json:"campaign_type,omitempty"`                      // 计划类型
	CampaignSubType                int               `json:"campaign_sub_type,omitempty"`                  // 计划子类型 商品库推广计划类型的老数据会返回【后续会进行下线】。4：DPA，5：SDPA
	AdType                         int               `json:"ad_type,omitempty"`                            // 广告计划类型 0:信息流，1:搜索
	BidType                        int               `json:"bid_type,omitempty"`                           // 出价类型
	AutoAdjust                     int               `json:"auto_adjust,omitempty"`                        // 自动调控开关 0:关闭，1:开启
	AutoBuild                      int               `json:"auto_build,omitempty"`                         // 自动基建开关 0:关闭，1:开启
	AutoBuildNameRule              AutoBuildNameRule `json:"auto_build_name_rule,omitempty"`               // 自动基建广告命名规则
	AutoManage                     int               `json:"auto_manage,omitempty"`                        // 智能投放开关 0:关闭，1:开启
	CampaignOCPXActionType         int               `json:"campaign_ocpx_action_type,omitempty"`          // 智能模式下的优化目标
	CampaignOCPXActionName         string            `json:"campaign_ocpx_action_type_name,omitempty"`     // 智能模式下的优化目标名称
	CampaignDeepConversionType     int               `json:"campaign_deep_conversion_type,omitempty"`      // 智能模式下的深度优化目标
	CampaignDeepConversionTypeName string            `json:"campaign_deep_conversion_type_name,omitempty"` // 智能模式下的深度优化目标名称
	CreateTime                     string            `json:"create_time,omitempty"`                        // 创建时间
	UpdateTime                     string            `json:"update_time,omitempty"`                        // 更新时间
}

// AutoBuildNameRule 自动基建命名规则
type AutoBuildNameRule struct {
	// UnitNameRule 广告组名称命名规则; 必须同时包含[日期]和[序号]宏变量，eg: 系统自动搭建_[日期][序号]
	UnitNameRule string `json:"unit_name_rule,omitempty"`
	// CreativeNameRule 广告创意名称命名规则; 必须同时包含[日期]和[序号]宏变量，eg: 系统自动搭建_[日期][序号]
	CreativeNameRule string `json:"creative_name_rule,omitempty"`
}
