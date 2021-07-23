package report

// MerchantStat 小店通转化数据
type MerchantStat struct {
	// ReportDate 日期
	ReportDate string `json:"report_date,omitempty"`
	// ReportHour 小时
	ReportHour int `json:"report_hour,omitempty"`
	// ReportDateHour 日期+小时
	ReportDateHour string `json:"report_date_hour,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignID 计划 ID
	CampaignID int64 `json:"campaign_id,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// UnitID 广告组 ID
	UnitID int64 `json:"unit_id,omitempty"`
	// CreativeName 创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// CreativeID 创意 id
	CreativeID int64 `json:"creative_id,omitempty"`
	// CostTotal 花费
	CostTotal int64 `json:"cost_total,omitempty"`
	// Impression 封面曝光数
	Impression int64 `json:"impression,omitempty"`
	// PhotoClick 封面点击数
	PhotoClick int64 `json:"photo_click,omitempty"`
	// PhotoClickRatio 封面点击率
	PhotoClickRatio float64 `json:"photo_click_ratio,omitempty"`
	// Click 素材曝光数
	Click int64 `json:"click,omitempty"`
	// ActionbarClick 行为数
	ActionbarClick int64 `json:"actionbar_click,omitempty"`
	// ActionRatio 行为率
	ActionRatio float64 `json:"action_ratio,omitempty"`
	// MerchantPhotoImpression1kCost 平均千次封面曝光花费(元)
	MerchantPhotoImpression1kCost float64 `json:"merchant_photo_impression_1k_cost,omitempty"`
	// MerchantPhotoClickCost 平均封面点击单价(元)
	MerchantPhotoClickCost float64 `json:"merchant_photo_click_cost,omitempty"`
	// MerchantImpression1kCost 平均千次素材曝光花费(元)
	MerchantImpression1kCost float64 `json:"merchant_impression_1k_cost,omitempty"`
	// MerchantClickCost 平均行为单价(元)
	MerchantClickCost float64 `json:"merchant_click_cost,omitempty"`
	// Play3sRatio 3s 播放率
	Play3sRatio float64 `json:"play_3s_ratio,omitempty"`
	Play5sRatio float64 `json:"play_5s_ratio,omitempty"`
	// Play5sRatio 5s 播放率
	// PlayEndRatio 完播率
	PlayEndRatio float64 `json:"play_end_ratio,omitempty"`
	// Share 分享数
	Share int64 `json:"share,omitempty"`
	// Comment 评论数
	Comment int64 `json:"comment,omitempty"`
	// Likes 点赞数
	Likes int64 `json:"likes,omitempty"`
	// Report 举报数
	Report int64 `json:"report,omitempty"`
	// Block 拉黑数
	Block int64 `json:"block,omitempty"`
	// Negative 减少此类作品数
	Negative int64 `json:"negative,omitempty"`
	// MerchantRecoFans 涨粉量
	MerchantRecoFans int64 `json:"merchant_reco_fans,omitempty"`
	// RecoFansCost 涨粉成本（元）
	RecoFansCost float64 `json:"reco_fans_cost,omitempty"`
	// PaidOrder 支付订单数
	PaidOrder int64 `json:"paid_order,omitempty"`
	// OrderCost 下单成本
	OrderCost float64 `json:"order_cost,omitempty"`
	// Gmv GMV
	Gmv float64 `json:"gmv,omitempty"`
	// T0Gmv 当日累计 GMV
	T0Gmv int64 `json:"t0_gmv,omitempty"`
	// T1Gmv 投后 1 日累计 GMV
	T1Gmv int64 `json:"t1_gmv,omitempty"`
	// T3Gmv 投后 3 日累计 GMV
	T3Gmv int64 `json:"t3_gmv,omitempty"`
	// T7Gmv 投后 7 日累计 GMV
	T7Gmv int64 `json:"t7_gmv,omitempty"`
	// T15Gmv 投后 15 GMV
	T15Gmv int64 `json:"t15_gmv,omitempty"`
	// T30Gmv 投后 30 GMV
	T30Gmv int64 `json:"t30_gmv,omitempty"`
	// Roi ROI
	Roi float64 `json:"roi,omitempty"`
	// T0Roi 当日累计 ROI
	T0Roi float64 `json:"t0Roi,omitempty"`
	// T1Roi 投后 1 日累计 ROI
	T1Roi float64 `json:"t1Roi,omitempty"`
	// T3Roi 投后 3 日累计 ROI
	T3Roi float64 `json:"t3Roi,omitempty"`
	// T7Roi 投后 7 日累计 ROI
	T7Roi float64 `json:"t7Roi,omitempty"`
	// T15Roi 投后 15 日累计 ROI
	T15Roi float64 `json:"t15Roi,omitempty"`
	// T30Roi 投后 30 日累计 ROI
	T30Roi float64 `json:"t30Roi,omitempty"`
	// T7Retention 投后 7 日涨粉留存量
	T7Retention int64 `json:"t7_retention,omitempty"`
	// T30Retention 投后 30 日涨粉留存量
	T30Retention int64 `json:"t30_retention,omitempty"`
	// T7RetentionRatio 投后 7 日涨粉留存率
	T7RetentionRatio float64 `json:"t7_retention_ratio,omitempty"`
	// T30RetentionRatio 投后 30 日涨粉留存率
	T30RetentionRatio float64 `json:"t30_retention_ratio,omitempty"`
	// T0OrderCnt 当日累计订单成交量
	T0OrderCnt int64 `json:"t0_order_cnt,omitempty"`
	// T1OrderCnt 投后 1 日累计订单成交量
	T1OrderCnt int64 `json:"t1_order_cnt,omitempty"`
	// T30OrderCnt 投后 3 日累计订单成交量
	T3OrderCnt int64 `json:"t3_order_cnt,omitempty"`
	// T7OrderCnt 投后 7 日累计订单成交量
	T7OrderCnt int64 `json:"t7_order_cnt,omitempty"`
	// T15OrderCnt 投后 15 日累计订单成交量
	T15OrderCnt int64 `json:"t15_order_cnt,omitempty"`
	// T30OrderCnt 投后 30 日累计订单成交量
	T30OrderCnt int64 `json:"t30_order_cnt,omitempty"`
}
