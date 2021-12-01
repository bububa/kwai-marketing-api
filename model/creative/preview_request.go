package creative

import "encoding/json"

// PreviewRequest 创意体验 API Request
type PreviewRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// UserIDs 快手id; 一次不超过十个
	UserIDs []string `json:"user_ids,omitempty"`
	// CreativeID 创意id; unit_type为4时，必填
	CreativeID int64 `json:"creative_id,omitempty"`
	// UnitID 组id; unit_type为7时，必填
	UnitID int64 `json:"unit_id,omitempty"`
	// UnitType 组类型; 4：自定义创意 7：程序化创意2.0
	UnitType int `json:"unit_type,omitempty"`
	// Phones 用户要推送创意到哪个手机号; 原本是传userId，由于允许邮箱开户后，不一定每个userId都有手机号，所以改成传手机号
	Phones []string `json:"phones,omitempty"`
}

// Url implement PostRequest interface
func (r PreviewRequest) Url() string {
	return "v1/creative/preview"
}

// Encode implement PostRequest interface
func (r PreviewRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
