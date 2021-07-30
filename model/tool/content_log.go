package tool

// ContentLog 日志内容
type ContentLog struct {
	// FieldName 字段名称
	FieldName string `json:"field_name,omitempty"`
	// OriginalData 原始数据
	OriginalData string `json:"original_data,omitempty"`
	// UpdateData 修改数据
	UpdateData string `json:"update_data,omitempty"`
}
