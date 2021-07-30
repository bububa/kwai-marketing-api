package dmp

import "encoding/json"

// PopulationAccountsPushRequest 人群包跨账户推送 APIRequest
type PopulationAccountsPushRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// OrientationID 人群包ID
	OrientationID int64 `json:"orientation_id,omitempty"`
	// DestAccountIDs 要推送的账户ids
	DestAccountIDs []int64 `json:"dest_account_ids,omitempty"`
}

// Url implement PostRequest interface
func (r PopulationAccountsPushRequest) Url() string {
	return "v1/dmp/population/accounts/push"
}

// Encode implement PostRequest interface
func (r PopulationAccountsPushRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
