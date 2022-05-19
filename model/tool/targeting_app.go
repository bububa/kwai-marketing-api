package tool

// TargetingApp 应用定向
type TargetingApp struct {
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
}
