package preput

import "encoding/json"

// PredicationTaskManagementRequest 投前预估任务管理接口 API Request
type PredicationTaskManagementRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ManagementType 管理类型
	MamagementType int `json:"management_type,omitempty"`
	// PredicationIDs 投前任务ids
	PredicationIDs []uint64 `json:"predication_ids,omitempty"`
}

// Url implement PostRequest interface
func (r PredicationTaskManagementRequest) Url() string {
	return "/gw/dsp/v1/pre_put/predication_task/management"
}

// Encode implement PostRequest interface
func (r PredicationTaskManagementRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
