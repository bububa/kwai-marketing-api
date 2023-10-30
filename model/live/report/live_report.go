package report

type ListLiveReportResponse struct {
	TotalCount int          `json:"total_count"` // 数据的总行数
	Details    []LiveReport `json:"details"`     // 数据明细信息
	StatDate   string       `json:"stat_date"`   // 数据日期，格式：YYYY-MM-DD
	StatHour   int          `json:"stat_hour"`   // 数据小时
}

// 主播统计数据结构体
type LiveReport struct {
	LiveStartTime                 int64   `json:"live_start_time"`
	LiveEndTime                   int64   `json:"live_end_time"`
	AdLiveFollow                  int64   `json:"ad_live_follow"`                    // 粉丝关注数
	LivePlayedStarted             float64 `json:"live_played_started"`               // 直播间进人数
	LivePlayedStartedCost         float64 `json:"live_played_started_cost"`          // 直播间进入成本
	StandardLivePlayedStarted     int     `json:"standard_live_played_started"`		 // 标准直播间开始播放数
	LiveRoomAvgPlayedSeconds      float64 `json:"live_room_avg_played_seconds"`      // 直播间平均观看时长
	ConversionComponentRate       float64 `json:"conversion_component_rate"`         // 组件点击率
	ConversionComponentClick      int64   `json:"conversion_component_click"`        // 组件点击量
	AdAppDownloadHalfImpression   int64   `json:"ad_app_download_half_impression"`   // 安卓APP下载类半屏落地页曝光
	UserId                        int64   `json:"user_id"`							// 主播ID
	AdLandingPageImpression       int64   `json:"ad_landing_page_impression"`        // 广告主第三方半屏落地页曝光
	SimpleLivePlayedStarted       int     `json:"simple_live_played_started"` //  简易直播间开始播放数
	LiveStreamId                  int64   `json:"live_stream_id"`
	AdLiveShare                   int64   `json:"ad_live_share"`                     // 直播间分享数
	AdLiveComment                 int64   `json:"ad_live_comment"`                   // 直播间评论数
	AdLiveFollowCost              float64 `json:"ad_live_follow_cost"`               // 粉丝关注成本
	ConversionComponentImpression int64   `json:"conversion_component_impression"`   // 组件展示量
}

