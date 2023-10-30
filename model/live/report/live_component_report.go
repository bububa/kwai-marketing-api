package report

import "encoding/json"

type ListLiveComponentReportRequest struct {
	AdvertiserID        uint64 `json:"advertiser_id"`        // 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回
	JingleBellID        int64  `json:"jingle_bell_id"`       // 小铃铛ID
	StartDateMin        string `json:"start_date_min"`       // 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	EndDateMin          string `json:"end_date_min"`         // 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	StartDate           string `json:"start_date"`           // 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明
	EndDate             string `json:"end_date"`             // 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明
	TemporalGranularity string `json:"temporal_granularity"` // 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据
	Page                int    `json:"page"`                 // 请求的页码，默认为 1
	PageSize            int    `json:"page_size"`            // 每页行数，默认为 20，最大支持 2000
}

// Url implement PostRequest interface
func (r ListLiveComponentReportRequest) Url() string {
	return "gw/dsp/v1/report/live_component_report"
}

// Encode implement PostRequest interface
func (r ListLiveComponentReportRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type ListLiveComponentReportResponse struct {
	TotalCount int                   `json:"total_count"` // 数据的总行数
	Details    []LiveComponentReport `json:"details"`     // 数据明细信息
}

type LiveComponentReport struct {
	Charge                                     float64 `json:"charge"`
	Show                                       int     `json:"show"`
	Aclick                                     int     `json:"aclick"`
	Bclick                                     int     `json:"bclick"`
	Share                                      int     `json:"share"`
	Comment                                    int     `json:"comment"`
	Like                                       int     `json:"like"`
	Follow                                     int     `json:"follow"`
	Report                                     int     `json:"report"`
	Block                                      int     `json:"block"`
	Negative                                   int     `json:"negative"`
	Activation                                 int     `json:"activation"`
	Submit                                     int     `json:"submit"`
	AdScene                                    string  `json:"adScene"`
	AdScene1                                   string  `json:"ad_scene"`
	PlacementType                              string  `json:"placement_type"`
	CancelFollow                               int     `json:"cancel_follow"`
	DownloadStarted                            int     `json:"download_started"`
	DownloadCompleted                          int     `json:"download_completed"`
	PhotoClick                                 int     `json:"photo_click"`
	PhotoClickRatio                            float64 `json:"photo_click_ratio"`
	ActionRatio                                float64 `json:"action_ratio"`
	Impression1KCost                           float64 `json:"impression_1k_cost"`
	PhotoClickCost                             float64 `json:"photo_click_cost"`
	Click1KCost                                float64 `json:"click_1k_cost"`
	ActionCost                                 float64 `json:"action_cost"`
	EventPayFirstDay                           int     `json:"event_pay_first_day"`
	EventPayPurchaseAmountFirstDay             float64 `json:"event_pay_purchase_amount_first_day"`
	EventPayFirstDayRoi                        float64 `json:"event_pay_first_day_roi"`
	EventPay                                   int     `json:"event_pay"`
	EventPayPurchaseAmount                     float64 `json:"event_pay_purchase_amount"`
	EventPayPurchaseAmount30DayByConversion    float64 `json:"event_pay_purchase_amount_30_day_by_conversion"`
	EventPayPurchaseAmount30DayByConversionRoi float64 `json:"event_pay_purchase_amount_30_day_by_conversion_roi"`
	EventPayRoi                                float64 `json:"event_pay_roi"`
	EventRegister                              int     `json:"event_register"`
	EventRegisterCost                          float64 `json:"event_register_cost"`
	EventRegisterRatio                         float64 `json:"event_register_ratio"`
	EventJinJianApp                            int     `json:"event_jin_jian_app"`
	EventJinJianAppCost                        float64 `json:"event_jin_jian_app_cost"`
	EventCreditGrantApp                        int     `json:"event_credit_grant_app"`
	EventCreditGrantAppCost                    float64 `json:"event_credit_grant_app_cost"`
	EventCreditGrantAppRatio                   float64 `json:"event_credit_grant_app_ratio"`
	EventOrderPaid                             int     `json:"event_order_paid"`
	EventOrderPaidPurchaseAmount               float64 `json:"event_order_paid_purchase_amount"`
	EventOrderPaidCost                         float64 `json:"event_order_paid_cost"`
	FormCount                                  int     `json:"form_count"`
	FormCost                                   float64 `json:"form_cost"`
	FormActionRatio                            float64 `json:"form_action_ratio"`
	EventJinJianLandingPage                    int     `json:"event_jin_jian_landing_page"`
	EventJinJianLandingPageCost                float64 `json:"event_jin_jian_landing_page_cost"`
	EventCreditGrantLandingPage                int     `json:"event_credit_grant_landing_page"`
	EventCreditGrantLandingPageCost            float64 `json:"event_credit_grant_landing_page_cost"`
	EventCreditGrantLandingRatio               float64 `json:"event_credit_grant_landing_ratio"`
	EventNextDayStayCost                       float64 `json:"event_next_day_stay_cost"`
	EventNextDayStayRatio                      float64 `json:"event_next_day_stay_ratio"`
	EventNextDayStay                           int     `json:"event_next_day_stay"`
	Play3SRatio                                float64 `json:"play_3s_ratio"`
	EventValidClues                            int     `json:"event_valid_clues"`
	EventValidCluesCost                        float64 `json:"event_valid_clues_cost"`
	AdProductCnt                               int     `json:"ad_product_cnt"`
	EventGoodsView                             int     `json:"event_goods_view"`
	MerchantRecoFans                           int     `json:"merchant_reco_fans"`
	EventOrderAmountRoi                        float64 `json:"event_order_amount_roi"`
	EventGoodsViewCost                         float64 `json:"event_goods_view_cost"`
	MerchantRecoFansCost                       float64 `json:"merchant_reco_fans_cost"`
	EventNewUserPay                            int     `json:"event_new_user_pay"`
	EventNewUserPayCost                        float64 `json:"event_new_user_pay_cost"`
	EventNewUserPayRatio                       float64 `json:"event_new_user_pay_ratio"`
	EventButtonClick                           int     `json:"event_button_click"`
	EventButtonClickCost                       float64 `json:"event_button_click_cost"`
	EventButtonClickRatio                      float64 `json:"event_button_click_ratio"`
	Play5SRatio                                float64 `json:"play_5s_ratio"`
	PlayEndRatio                               float64 `json:"play_end_ratio"`
	EventOrderPaidRoi                          float64 `json:"event_order_paid_roi"`
	EventNewUserJinjianApp                     int     `json:"event_new_user_jinjian_app"`
	EventNewUserJinjianAppCost                 float64 `json:"event_new_user_jinjian_app_cost"`
	EventNewUserJinjianAppRoi                  float64 `json:"event_new_user_jinjian_app_roi"`
	EventNewUserCreditGrantApp                 int     `json:"event_new_user_credit_grant_app"`
	EventNewUserCreditGrantAppCost             float64 `json:"event_new_user_credit_grant_app_cost"`
	EventNewUserCreditGrantAppRoi              float64 `json:"event_new_user_credit_grant_app_roi"`
	EventNewUserJinjianPage                    int     `json:"event_new_user_jinjian_page"`
	EventNewUserJinjianPageCost                float64 `json:"event_new_user_jinjian_page_cost"`
	EventNewUserJinjianPageRoi                 float64 `json:"event_new_user_jinjian_page_roi"`
	EventNewUserCreditGrantPage                int     `json:"event_new_user_credit_grant_page"`
	EventNewUserCreditGrantPageCost            float64 `json:"event_new_user_credit_grant_page_cost"`
	EventNewUserCreditGrantPageRoi             float64 `json:"event_new_user_credit_grant_page_roi"`
	EventAppointForm                           int     `json:"event_appoint_form"`
	EventAppointFormCost                       float64 `json:"event_appoint_form_cost"`
	EventAppointFormRatio                      float64 `json:"event_appoint_form_ratio"`
	EventAppointJumpClick                      int     `json:"event_appoint_jump_click"`
	EventAppointJumpClickCost                  float64 `json:"event_appoint_jump_click_cost"`
	EventAppointJumpClickRatio                 float64 `json:"event_appoint_jump_click_ratio"`
	EventNextDayStayNew                        int     `json:"event_next_day_stay_new"`
	EventNextDayStayNewCost                    float64 `json:"event_next_day_stay_new_cost"`
	EventNextDayStayNewRatio                   float64 `json:"event_next_day_stay_new_ratio"`
	EventMultiPaySevenDayByConversion          float64 `json:"event_multi_pay_seven_day_by_conversion"`
	EventMultiPaySevenDayByConversionCost      float64 `json:"event_multi_pay_seven_day_by_conversion_cost"`
	LiveRoomAvgPlayedSeconds                   float64 `json:"live_room_avg_played_seconds"`
	AdLiveShare                                int     `json:"ad_live_share"`
	AdLiveComment                              int     `json:"ad_live_comment"`
	LivePlayedStarted                          float64 `json:"live_played_started"`
	LivePlayedStartedCost                      float64 `json:"live_played_started_cost"`
	AdLiveFollow                               int     `json:"ad_live_follow"`
	AdLiveFollowCost                           float64 `json:"ad_live_follow_cost"`
	SimpleLivePlayedStarted                    int     `json:"simple_live_played_started"`
	StandardLivePlayedStarted                  int     `json:"standard_live_played_started"`
	ConversionComponentImpression              int     `json:"conversion_component_impression"`
	ConversionComponentClick                   int     `json:"conversion_component_click"`
	ConversionComponentRate                    float64 `json:"conversion_component_rate"`
	AdLandingPageImpression                    int     `json:"ad_landing_page_impression"`
	AdAppDownloadHalfImpression                int     `json:"ad_app_download_half_impression"`
	EventDrawCreditLine                        int     `json:"event_draw_credit_line"`
	JingleBellId                               int64   `json:"jingle_bell_id"`
}
