package dmp

// Population 人群包
type Population struct {
	// OrientationID 人群包ID
	OrientationID uint64 `json:"orientation_id,omitempty"`
	// OrientationName 人群包名称
	OrientationName string `json:"orientation_name,omitempty"`
	// Type 人群数据类型; 1：IMEI；2：IDFA；3：IMEI_MD5；4：IDFA_MD5；5：手机号-MD5；7：OAID；8：OAID_MD5
	Type int `json:"type,omitempty"`
	// PopulationType 人群包类型; 1, "上传人群"；2, "广告人群"；3, "主题专区"；4, "逻辑规则"；5, "人群扩展"；6, "平台定制"；7, "定制付费"；8, "网红粉丝类别"；9, "内容付费行为"；10, "移动应用安装"；11, "快手使用活跃度"；12, "行业分类"；13, "商业兴趣"；14, "固化标签"；15, "行业偏好"；16, "第三方标签"；17, "产品关键词"；19, "应用渗透率"；22, "指定网红"；23, "行业分类"（同12）
	PopulationType int `json:"population_type,omitempty"`
	// RecordSize 上传数量
	RecordSize int64 `json:"record_size,omitempty"`
	// MatchSize 匹配数量
	MatchSize int64 `json:"match_size,omitempty"`
	// CoverNum 覆盖人群; 返回值固定为0
	CoverNum int64 `json:"cover_num,omitempty"`
	// Status 人群包状态; 0, "计算中"；1, "已生效"；2, "已删除"；3, "上线中"；4, "已上线"；5, "计算失败"；6, "上线失败"；7, "已失效"
	Status int `json:"status,omitempty"`
	// CreateTime 创建时间; 格式：13位毫秒级时间戳
	CreateTime int64 `json:"create_time,omitempty"`
	// ThirdPlatformCode 付费人群包-第3方平台code
	ThirdPlatformCode uint64 `json:"third_platform_code,omitempty"`
	// ThirdPlatformName 付费人群包第三方供应商名称
	ThirdPlatformName string `json:"third_platform_name,omitempty"`
}
