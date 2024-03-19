package unit

import "github.com/bububa/kwai-marketing-api/model"

// GetMonitorURLsRequest 批量获取监测链接接口 API Request
type GetMonitorURLsRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitIDs 广告组 ID 集
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
}

// Url implement PostRequest interface
func (r GetMonitorURLsRequest) Url() string {
	return "gw/dsp/v3/unit/getMonitorUrls"
}

// Encode implement PostRequest interface
func (r GetMonitorURLsRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// GetMonitorURLsResponse 批量获取监测链接接口 API Response
type GetMonitorURLsResponse struct {
	UnitMonitorURLs []UnitMonitorURL `json:"unit_monitor_urls,omitempty"`
}

// BatchUpdateMonitorURLsRequest 监测链接批量更新接口 API Request
type BatchUpdateMonitorURLsRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitMonitorURLs 监测链接详情
	UnitMonitorURLs []UnitMonitorURL `json:"unit_monitor_urls,omitempty"`
}

// Url implement PostRequest interface
func (r BatchUpdateMonitorURLsRequest) Url() string {
	return "gw/dsp/v3/unit/batchUpdateMonitorUrls"
}

// Encode implement PostRequest interface
func (r BatchUpdateMonitorURLsRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UnitMonitorURL 监测链接
type UnitMonitorURL struct {
	// UnitID 广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// IsDpa 是否是DPA（DPA的监测链接在广告组上）
	IsDpa bool `json:"is_dpa,omitempty"`
	// ExistValidCreative 是否存在有效创意
	ExistValidCreative bool `json:"exist_valid_creative,omitempty"`
	// ActionbarClickURL actionBar点击监测，命中有效触点白名单表示有效触点监测链接（限：快手主站/极速版场景）
	ActionbarClickURL string `json:"actionbar_click_url,omitempty"`
	// ClickURL 点击监测（排除粉丝直播推广）
	ClickURL string `json:"click_url,omitempty"`
	// ImpressionURL 展示监测
	ImpressionURL string `json:"impression_url,omitempty"`
	// LiveTrackURL 粉丝直播推广点击监测
	LiveTrackURL string `json:"live_track_url,omitempty"`
	// AdPhotoPlayedT3sURL 粉丝直播推广曝光监测
	AdPhotoPlayedT3sURL string `json:"ad_photo_played_t3s_url,omitempty"`
	// Result 修改是否成功
	Result bool `json:"result,omitempty"`
	// Message 修改成功与否提示信息
	Message string `json:"message,omitempty"`
}
