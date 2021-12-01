package creative

// DisplayInfo 广告展示信息
type DisplayInfo struct {
	// Description 广告语
	Description string `json:"description,omitempty"`
	// ActionBarText 行动号召按钮文案
	ActionBarText string `json:"action_bar_text,omitempty"`
}
