package dmp

import "encoding/json"

// PopulationDeleteRequest 人群包删除接口APIRequest
type PopulationDeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OrientationID 人群包ID
	OrientationID uint64 `json:"orientation_id,omitempty"`
}

// Url implement PostRequest interface
func (r PopulationDeleteRequest) Url() string {
	return "v1/dmp/population/delete"
}

// Encode implement PostRequest interface
func (r PopulationDeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
