package unit

// DiverseData 应用信息
type DiverseData struct {
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// AppPackageName 应用包名
	AppPackageName string `json:"app_package_name,omitempty"`
	// DeviceOsType 应用操作系统类型; 0：未知，1：ANDROID，2：IO
	DeviceOsType int `json:"device_os_type,omitempty"`
}
