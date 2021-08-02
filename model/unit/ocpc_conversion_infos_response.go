package unit

// OcpcConversionInfosResponse 获取可选的深度转化目标 API Response
type OcpcConversionInfosResponse struct {
	// IsActivate 是否有激活出价权限
	IsActivate int `json:"is_activate,omitempty"`
	// IsFormSubmit 是否有表单出价权限
	IsFormSubmit int `json:"is_form_submit,omitempty"`
	// IsCreditGrant 是否有授信出价权限
	IsCreditGrant int `json:"is_credit_grant,omitempty"`
	// IsPurchase 是否有付费出价权限
	IsPurchase int `json:"is_purchase,omitempty"`
	// IsWanJian 是否有完件出价权限
	IsWanJian int `json:"is_wan_jian,omitempty"`
	// IsValidClue 是否可以使用有效线索
	IsValidClue int `json:"is_valid_clue,omitempty"`
	// IsFirstdayRoi 是否可以使用首日ROI
	IsFirstdayRoi int `json:"is_firstday_roi,omitempty"`
	// IsRegister 是否可以使用注册优化目标
	IsRegister int `json:"is_register,omitempty"`
	// IsOrderSubmit 是否可以使用订单提交优化目标
	IsOrderSubmit int `json:"is_order_submit,omitempty"`
	// IsAddWechat 是否可用微信复制
	IsAddWechat int `json:"is_add_wechat,omitempty"`
	// IsAppInvoked 是否可以使用唤起应用优化目标
	IsAppInvoked int `json:"is_app_invoked,omitempty"`
	// IsMultiConversion 是否可用多转化事件
	IsMultiConversion int `json:"is_multi_conversion,omitempty"`
	// IsAdWatchTimes 是否可用广告观看次数
	IsAdWatchTimes int `json:"is_ad_watch_times,omitempty"`
	// IsOrderPaid 是否可以使用订单支付转化目标
	IsOrderPaid int `json:"is_order_paid,omitempty"`
	// IsSevenDayRoi 是否可以使用7日ROI目标
	IsSevenDayRoi int `json:"is_seven_day_roi,omitempty"`
	// IsKeyAction 是否可以使用关键行为目标
	IsKeyAction int `json:"is_key_action,omitempty"`
	// DeepConversionTypes 深度转化类型
	DeepConversionTypes []DeepConversionType `json:"deep_conversion_types,omitempty"`
}

// DeepConversionType 深度转化类型
type DeepConversionType struct {
	// Desc 深度转化类型描述
	Desc string `json:"desc,omitempty"`
	// DeepConversionType 深度转化类型
	DeepConversionType int `json:"deep_conversion_type,omitempty"`
}
