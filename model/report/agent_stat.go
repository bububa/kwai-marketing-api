package report

type AgentStat struct {
	DateTime                        string `json:"date_time,omitempty"`                              // 数据日期，格式：YYYY-MM-DD
	AccountID                       int64  `json:"account_id,omitempty"`                             // 广告主ID
	UserID                          int64  `json:"user_id,omitempty"`                                // 快手id
	AccountName                     string `json:"account_name,omitempty"`                           // 广告主名称
	TotalChargedInYuan              int64  `json:"total_charged_in_yuan,omitempty"`                  // 总消耗（元）
	TotalBalanceInYuan              int64  `json:"total_balance_in_yuan,omitempty"`                  // 总余额（元）
	RealChargedInYuan               int64  `json:"real_charged_in_yuan,omitempty"`                   // 现金消耗（元）
	TotalRebateRechargedInYuan      int64  `json:"total_real_recharged_in_yuan,omitempty"`           // 返点消耗（元）
	ContractRebateRealChargedInYuan int64  `json:"contract_rebate_real_recharged_in_yuan,omitempty"` // 框返消耗（元）
	DirectRebateRealChargedInYuan   int64  `json:"direct_rebate_real_charged_in_yuan,omitempty"`     // 激励消耗（元）
	CreditRealChargedInYuan         int64  `json:"credit_real_charged_in_yuan,omitempty"`            // 信用消耗（元）
	ChargeDayOnDayPercent           string `json:"charge_day_on_day_percent,omitempty"`              // 消耗环比
	AdPhotoImpression               int64  `json:"ad_photo_impression,omitempty"`                    // 封面曝光数
	AdPhotoClick                    int64  `json:"ad_photo_click,omitempty"`                         // 封面点击数
	AdItemClick                     int64  `json:"ad_item_click,omitempty"`                          // 行为数
	PhotoClickRatio                 string `json:"photo_click_ratio,omitempty"`                      // 封面点击率
	ItemClickRatio                  string `json:"item_click_ratio,omitempty"`                       // 行为点击率
	ChargedCampaignCount            int64  `json:"charged_campaign_count,omitempty"`                 // 有消费计划数
	ProductName                     string `json:"product_name,omitempty"`                           // 产品名
	CorporationName                 string `json:"corporation_name,omitempty"`                       // 企业名称
	FirstCostDay                    int64  `json:"first_cost_day,omitempty"`                         // 首次消耗日期
	Industry                        string `json:"industry,omitempty"`                               // 一级行业
	SecondIndustry                  string `json:"second_industry,omitempty"`                        // 二级行业
}
