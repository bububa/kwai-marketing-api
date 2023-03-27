package page

// Page 魔力建站落地页
type Page struct {
	// ID 落地页 ID
	ID int64 `json:"id,omitempty"`
	// Comps 落地页包含的组件列表
	Comps []Component `json:"comps,omitempty"`
	// URL 落地页 URL
	URL string `json:"url,omitempty"`
	// FictionID 落地页绑定的小说 ID
	FictionID int64 `json:"fiction_id,omitempty"`
	// BizType 落地页类型，0：站内；1：联盟；2：站内&联盟通投
	BizType int `json:"biz_type,omitempty"`
	// CreateTime 创建时间，格式：yyyy-MM-dd hh:MM:ss
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 修改时间，格式：yyyy-MM-dd hh:MM:ss
	UpdateTime string `json:"update_time,omitempty"`
	// ConversionTypes 落地页包含的转化组件类型:APP_DOWNLOAD_ANDROID_DRAG 应用下载-安卓、APP_DOWNLOAD_IOS_DRAG
	// 应用下载-ios、BUTTON_DRAG 按钮、XIANSUO_FORM_DRAG 表单、WEI_XIN_DRAG 微信、CUSTOMER_SERVICE_DRAG 客服咨询、
	// COUPON_CARD 卡劵、TEL_DRAG 电话、WECHAT_GAME 小游戏、落地页组限制落地页必须含有相同转化组件类型
	ConversionTypes []string `json:"conversion_type,omitempty"`
	// Details JSON 返回值
	Details interface{} `json:"details,omitempty"`
}

// Component 落地页包含的组件
type Component struct {
	// ID 组件 ID
	ID int64 `json:"id,omitempty"`
	// Type 组件类型，0: 图片；1: 文本；2: 表单；3: 按钮；4: 轮播图；5: 视频；6: 地图；7: 应用下载；16: 空白组件；34: 小游戏；
	Type int `json:"type,omitempty"`
	// Name 组件名称
	Name string `json:"name,omitempty"`
	// Props 组件属性，仅当 Type 为 7 时有用, 其他类型没有这个属性
	Props map[string]interface{} `json:"props,omitempty"`
	// WechatGameID 小游戏类型对应的 ID，Type 为 34 时有用
	WechatGameID int64 `json:"wechat_game_id,omitempty"`
	// ButtonText 小游戏类型对应的按钮文案，Type 为 34 时有用
	ButtonText string `json:"button_text,omitempty"`
	// GameName 小游戏的名称，Type 为 34 时有用
	GameName string `json:"game_name,omitempty"`
	// Description 小游戏的说明，Type 为 34 时有用
	Description string `json:"description,omitempty"`
}
