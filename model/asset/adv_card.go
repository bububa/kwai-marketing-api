package asset

// AdvCard 高级创意
type AdvCard struct {
	// AdvCardID 卡片 id
	AdvCardID uint64 `json:"adv_card_id,omitempty"`
	// TemplateName 模版名
	TemplateName string `json:"template_name,omitempty"`
	// UnitCount 关联广告组数
	UnitCount int `json:"unit_count,omitempty"`
	// URL 图片 url
	URL string `json:"url,omitempty"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// SubTitle 副标题
	SubTitle string `json:"sub_title,omitempty"`
	// Price 原价格(单位：分)
	Price int64 `json:"price,omitempty"`
	// SalePrice 售卖价(单位：分)
	SalePrice int64 `json:"sale_price,omitempty"`
	// CardType 卡片类型; 100:图片卡片 101:多利益卡-图文 102：多利益卡-多标签 103：电商促销样式 104：快捷评论卡
	CardType int `json:"card_type,omitempty"`
	// ContentType 卡片内容类型; 当 card_type = 104 必填，当创建 emoji 快捷评论卡时，填 2
	ContentType int `json:"content_type,omitempty"`
	// EmojiList emoji 信息; 当 card_type = 104 必填
	EmojiList []Emoji `json:"emoji_list,omitempty"`
}
