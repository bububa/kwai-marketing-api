package advertiser

// WhiteListResponse 获取可选白名单接口 API Response
type WhiteListResponse struct {
	// CreativeCategorySwitch 账户能否使用创意标签分类
	CreativeCategorySwitch int `json:"creative_category_switch,omitempty"`
	// ActionbarClickUrlSwitch 点击监测白名单，该用户是否在二跳白名单中，是否支持actionbar_click_url
	ActionbarClickUrlSwitch int `json:"actionbar_click_url_switch,omitempty"`
	// AdPhotoPlayedT3sUrlSwitch 曝光三秒白名单，此账户是否支持ad_photo_played_t3s_url
	AdPhotoPlayedT3sUrlSwitch int `json:"ad_photo_played_t3s_url_switch,omitempty"`
}
