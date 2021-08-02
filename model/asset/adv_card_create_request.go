package asset

import "encoding/json"

// AdvCardCreateRequest 创建高级创意接口API Request
type AdvCardCreateRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID int64 `json:'advertiser_id,omitempty'`
	// AdvList 卡片
	AdvList []AdvCard `json:"adv_list,omitempty"`
}

// Url implement PostRequest interface
func (r AdvCardCreateRequest) Url() string {
	return "v1/asset/adv_card/create"
}

// Encode implenment PostRequest interface
func (r AdvCardCreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
