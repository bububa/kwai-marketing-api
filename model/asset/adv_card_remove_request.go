package asset

import "encoding/json"

// AdvCardRemoveRequest 删除高级创意接口
type AdvCardRemoveRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// AdvCardID 卡片 id
	AdvCardID int64 `json:"adv_card_id,omitempty"`
}

// Url implement PostRequest interface
func (r AdvCardRemoveRequest) Url() string {
	return "v1/asset/adv_card/remove"
}

// Encode implement PostRequest interface
func (r AdvCardRemoveRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
