package lp

// Consult 咨询组件
type Consult struct {
	// ID 唯一 id，unit 创建参数 consult_id
	ID int64 `json:"id,omitempty"`
	// Title 客服名称
	Title string `json:"title,omitempty"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time,omitempty"`
	// LastSessionTime 最后会话时间
	LastSessionTime int64 `json:"last_session_time,omitempty"`
	// ThirdPartyAppID 第三方客服 app_id; 当前支持的枚举值如下：快聊客服：10000；美洽客服：10001；53 客服：10002；快商通客服：10003；易聊客服：10004；螳螂客服：10005
	ThirdPartyAppID int `json:"third_party_app_id,omitempty"`
	// ConsultantsNum 咨询人数; 历史累计数值
	ConsultantsNum int64 `json:"consultants_num,omitempty"`
}
