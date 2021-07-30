package unit

// Age 自定义年龄段
type Age struct {
	// Min 年龄最小限制
	Min int `json:"min,omitempty"`
	// Max 年龄最大限制
	Max int `json:"max,omitempty"`
}
