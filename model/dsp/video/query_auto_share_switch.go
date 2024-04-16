package video

import "github.com/bububa/kwai-marketing-api/model"

// QueryAutoShareSwitchRequest 查询账户共享视频库按钮是否开启 API Request
type QueryAutoShareSwitchRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement PostRequest interface
func (r QueryAutoShareSwitchRequest) Url() string {
	return "gw/dsp/video/queryAutoShareSwitch"
}

// Encode implement PostRequest interface
func (r QueryAutoShareSwitchRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// QueryAutoShareSwitchResponse 查询账户共享视频库按钮是否开启 API Response
type QueryAutoShareSwitchResponse struct {
	// SwitchStatus 开关状态：true代表开启，false代表关闭
	SwitchStatus bool `json:"switch_status,omitempty"`
	// UserID 所属的user_id
	UserID uint64 `json:"user_id,omitempty"`
}
