package jingbell

import "github.com/bububa/kwai-marketing-api/model"

// ShareRequest 小铃铛推送 API Request
type ShareRequest struct {
	// AdvertiserID 广告主userId
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LiveUserID 主播ID
	LiveUserID string `json:"live_user_id,omitempty"`
	// JingBellUserID 小铃铛ID
	JingBellUserID string `json:"jing_bell_user_id,omitempty"`
	// TargetAccountIDs 推送目标 accountIds
	TargetAccountIDs []string `json:"target_account_ids,omitempty"`
}

// Url implement PostRequest interface
func (r ShareRequest) Url() string {
	return "gw/dsp/jingBell/share"
}

// Encode implement PostRequest interface
func (r ShareRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ShareResponse 小铃铛推送 API Response
type ShareResponse struct {
	// Result 1成功 非1失败
	Result int `json:"result,omitempty"`
	// Data 如果失败，则是失败信息
	Data string `json:"data,omitempty"`
}
