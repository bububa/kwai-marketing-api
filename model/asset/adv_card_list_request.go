package asset

import "encoding/json"

// AdvCardListRequest 获取高级创意列表
type AdvCardListRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CardType 卡片类型; 100:图片卡片 101:多利益卡-图文 102：多利益卡-多标签 103：电商促销样式
	CardType int `json:"card_type,omitempty"`
	// Page 查询的页码数
	Page int `json:"page,omitempty"`
	// PageSize 单页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r AdvCardListRequest) Url() string {
	return "v1/asset/adv_card/list"
}

// Encode implement PostRequest interface
func (r AdvCardListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
