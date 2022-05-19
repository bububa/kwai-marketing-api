package tool

// Convert 转化目标
type Convert struct {
	// ConvertID 转化目标ID
	ConvertID int64 `json:"convert_id,omitempty"`
	// ConvertName 转化目标名称
	ConvertName string `json:"convert_name,omitempty"`
	// Type 转化目标类型; 1：JS布玛 2：Xpath 3：应用-SDK 7：应用-API
	Type int `json:"type,omitempty"`
	// ConvertTarget 转化目标; 1:表单提交 2:激活
	ConvertTarget int `json:"convert_target,omitempty"`
	// ConvertCount 转化目标次数
	ConvertCount int64 `json:"convert_count,omitempty"`
	// DeepConversionType 深度转化目标; 0：无 3：付费 7：次日留存 10：完件 11：授信 13：添加购物车 14：提交订单
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
	// DeepConversionCount 深度转化目标次数
	DeepConversionCount int64 `json:"deep_conversion_count,omitempty"`
	// UpdateTime 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// AppID 应用id; type为3、7时返回
	AppID uint64 `json:"app_id,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// ClickUrl 第三方点击按钮监测链接
	ClickUrl string `json:"click_url,omitempty"`
}
