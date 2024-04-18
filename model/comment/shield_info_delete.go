package comment

import "encoding/json"

// ShieldInfoDeleteRequest 删除屏蔽评论信息 API Request
type ShieldInfoDeleteRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ShieldInfoIDList 屏蔽信息id
	ShieldInfoIDList []uint64 `json:"shield_info_id_list,omitempty"`
}

// Url implement PostRequest interface
func (r ShieldInfoDeleteRequest) Url() string {
	return "v1/comment/shield_info/delete"
}

// Encode implement PostRequest interface
func (r ShieldInfoDeleteRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
