package dmp

import "encoding/json"

// PopulationPushRequest 人群包上线接口APIRequest
type PopulationPushRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OrientationID 人群包 ID; 只有 status 为 1 或 6 的人群包可上线，最多支持 4 个人群包状态同时为“上线中”，
	OrientationID uint64 `json:"orientation_id,omitempty"`
}

// Url implement PostRequest interface
func (r PopulationPushRequest) Url() string {
	return "v1/dmp/population/push"
}

// Encode implement PostRequest interface
func (r PopulationPushRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
