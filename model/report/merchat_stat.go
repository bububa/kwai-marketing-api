package report

type MerchantStat struct {
	ReportDate                    string  `json:"report_date,omitempty"`                       // 日期
	ReportHour                    int     `json:"report_hour,omitempty"`                       // 小时
	ReportDateHour                string  `json:"report_date_hour,omitempty"`                  // 日期+小时
	CampaignName                  string  `json:"campaign_name,omitempty"`                     // 计划名称
	CampaignID                    int64   `json:"campaign_id,omitempty"`                       // 计划 ID
	UnitName                      string  `json:"unit_name,omitempty"`                         // 广告组名称
	UnitID                        int64   `json:"unit_id,omitempty"`                           // 广告组 ID
	CreativeName                  string  `json:"creative_name,omitempty"`                     // 创意名称
	CreativeID                    int64   `json:"creative_id,omitempty"`                       // 创意 id
	CostTotal                     int64   `json:"cost_total,omitempty"`                        // 花费
	Impression                    int64   `json:"impression,omitempty"`                        // 封面曝光数
	PhotoClick                    int64   `json:"photo_click,omitempty"`                       // 封面点击数
	PhotoClickRatio               float64 `json:"photo_click_ratio,omitempty"`                 // 封面点击率
	Click                         int64   `json:"click,omitempty"`                             // 素材曝光数
	ActionbarClick                int64   `json:"actionbar_click,omitempty"`                   // 行为数
	ActionRatio                   float64 `json:"action_ratio,omitempty"`                      // 行为率
	MerchantPhotoImpression1kCost float64 `json:"merchant_photo_impression_1k_cost,omitempty"` // 平均千次封面曝光花费(元)
	MerchantPhotoClickCost        float64 `json:"merchant_photo_click_cost,omitempty"`         // 平均封面点击单价(元)
	MerchantImpression1kCost      float64 `json:"merchant_impression_1k_cost,omitempty"`       // 平均千次素材曝光花费(元)
	MerchantClickCost             float64 `json:"merchant_click_cost,omitempty"`               // 平均行为单价(元)
	Play3sRatio                   float64 `json:"play_3s_ratio,omitempty"`                     // 3s 播放率
	Play5sRatio                   float64 `json:"play_5s_ratio,omitempty"`                     // 5s 播放率
	PlayEndRatio                  float64 `json:"play_end_ratio,omitempty"`                    // 完播率
	Share                         int64   `json:"share,omitempty"`                             // 分享数
	Comment                       int64   `json:"comment,omitempty"`                           // 评论数
	Likes                         int64   `json:"likes,omitempty"`                             // 点赞数
	Report                        int64   `json:"report,omitempty"`                            // 举报数
	Block                         int64   `json:"block,omitempty"`                             // 拉黑数
	Negative                      int64   `json:"negative,omitempty"`                          // 减少此类作品数
	MerchantRecoFans              int64   `json:"merchant_reco_fans,omitempty"`                // 涨粉量
	RecoFansCost                  float64 `json:"reco_fans_cost,omitempty"`                    // 涨粉成本（元）
	PaidOrder                     int64   `json:"paid_order,omitempty"`                        // 支付订单数
	OrderCost                     float64 `json:"order_cost,omitempty"`                        // 下单成本
	Gmv                           float64 `json:"gmv,omitempty"`                               // GMV
	T0Gmv                         int64   `json:"t0_gmv,omitempty"`                            // 当日累计 GMV
	T1Gmv                         int64   `json:"t1_gmv,omitempty"`                            // 投后 1 日累计 GMV
	T3Gmv                         int64   `json:"t3_gmv,omitempty"`                            // 投后 3 日累计 GMV
	T7Gmv                         int64   `json:"t7_gmv,omitempty"`                            // 投后 7 日累计 GMV
	T15Gmv                        int64   `json:"t15_gmv,omitempty"`                           // 投后 15 GMV
	T30Gmv                        int64   `json:"t30_gmv,omitempty"`                           // 投后 30 GMV
	Roi                           float64 `json:"roi,omitempty"`                               // ROI
	T0Roi                         float64 `json:"t0Roi,omitempty"`                             // 当日累计 ROI
	T1Roi                         float64 `json:"t1Roi,omitempty"`                             // 投后 1 日累计 ROI
	T3Roi                         float64 `json:"t3Roi,omitempty"`                             // 投后 3 日累计 ROI
	T7Roi                         float64 `json:"t7Roi,omitempty"`                             // 投后 7 日累计 ROI
	T15Roi                        float64 `json:"t15Roi,omitempty"`                            // 投后 15 日累计 ROI
	T30Roi                        float64 `json:"t30Roi,omitempty"`                            // 投后 30 日累计 ROI
	T7Retention                   int64   `json:"t7_retention,omitempty"`                      // 投后 7 日涨粉留存量
	T30Retention                  int64   `json:"t30_retention,omitempty"`                     // 投后 30 日涨粉留存量
	T7RetentionRatio              float64 `json:"t7_retention_ratio,omitempty"`                // 投后 7 日涨粉留存率
	T30RetentionRatio             float64 `json:"t30_retention_ratio,omitempty"`               // 投后 30 日涨粉留存率
	T0OrderCnt                    int64   `json:"t0_order_cnt,omitempty"`                      // 当日累计订单成交量
	T1OrderCnt                    int64   `json:"t1_order_cnt,omitempty"`                      // 投后 1 日累计订单成交量
	T3OrderCnt                    int64   `json:"t3_order_cnt,omitempty"`                      // 投后 3 日累计订单成交量
	T7OrderCnt                    int64   `json:"t7_order_cnt,omitempty"`                      // 投后 7 日累计订单成交量
	T15OrderCnt                   int64   `json:"t15_order_cnt,omitempty"`                     // 投后 15 日累计订单成交量
	T30OrderCnt                   int64   `json:"t30_order_cnt,omitempty"`                     // 投后 30 日累计订单成交量
}
