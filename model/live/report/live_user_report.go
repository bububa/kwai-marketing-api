package report

import "encoding/json"

type ListLiveUserReportRequest struct {
	AdvertiserID        uint64 `json:"advertiser_id"`        // 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回
	StartDateMin        string `json:"start_date_min"`       // 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	EndDateMin          string `json:"end_date_min"`         // 增量拉取过滤筛选条件，格式【yyyy-MM-dd HH:mm】可选时间范围参见文档上方说明
	StartDate           string `json:"start_date"`           // 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明
	EndDate             string `json:"end_date"`             // 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明
	TemporalGranularity string `json:"temporal_granularity"` // 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据
	Page                int    `json:"page"`                 // 请求的页码，默认为 1
	PageSize            int    `json:"page_size"`            // 每页行数，默认为 20，最大支持 2000
}

// Url implement PostRequest interface
func (r ListLiveUserReportRequest) Url() string {
	return "gw/dsp/v1/report/live_user_report"
}

// Encode implement PostRequest interface
func (r ListLiveUserReportRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type ListLiveUserReportResponse struct {
	TotalCount int              `json:"total_count"` // 数据的总行数
	Details    []LiveUserReport `json:"details"`     // 数据明细信息
}

type LiveUserReport struct {
	LiveStartTime                 int64   `json:"live_start_time"`
	LiveEndTime                   int64   `json:"live_end_time"`
	AdLiveFollow                  int     `json:"ad_live_follow"`
	StandardLivePlayedStarted     int     `json:"standard_live_played_started"`
	LivePlayedStarted             float64 `json:"live_played_started"`
	LiveRoomAvgPlayedSeconds      float64 `json:"live_room_avg_played_seconds"`
	ConversionComponentRate       float64 `json:"conversion_component_rate"`
	LivePlayedStartedCost         float64 `json:"live_played_started_cost"`
	ConversionComponentClick      int     `json:"conversion_component_click"`
	AdAppDownloadHalfImpression   int     `json:"ad_app_download_half_impression"`
	UserId                        int64   `json:"user_id"`
	AdLandingPageImpression       int     `json:"ad_landing_page_impression"`
	SimpleLivePlayedStarted       int     `json:"simple_live_played_started"`
	LiveStreamId                  int64   `json:"live_stream_id"`
	AdLiveFollowCost              float64 `json:"ad_live_follow_cost"`
	AdLiveComment                 int     `json:"ad_live_comment"`
	ConversionComponentImpression int     `json:"conversion_component_impression"`
	AdLiveShare                   int     `json:"ad_live_share"`
}
