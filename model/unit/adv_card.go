package unit

// AdvCard 高级创意
type AdvCard struct {
	// ID 卡片id
	ID uint64 `json:"adv_card_id,omitempty"`
	// CardType 卡片类型; 100:图片卡片 101:多利益卡-图文 102：多利益卡-多标签 103：电商促销样式
	CardType int `json:"card_type,omitempty"`
	// Url 图片url
	Url string `json:"url,omitempty"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// SubTitle 副标题
	SubTitle string `json:"sub_title,omitempty"`
	// Price 原价格(单位：分)
	Price int64 `json:"price,omitempty"`
	// SalePrice 售卖价(单位：分)
	SalePrice int64 `json:"sale_price,omitempty"`
}
