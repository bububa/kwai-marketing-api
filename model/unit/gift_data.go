package unit

// GiftData 游戏礼包码
type GiftData struct {
	// TargetActionType 发送时机; android:30： 开始下载后 31：下载完成后 32：安装完成后 ios：2：检测到用户行为
	TargetActionType int `json:"target_action_type,omitempty"`
	// Code 礼包码; 最多20字符，只能字母+数字
	Code string `json:"code,omitempty"`
}
