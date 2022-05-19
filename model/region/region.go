package region

// Region 地域定向
type Region struct {
	// ID 地域ID
	ID uint64 `json:"id,omitempty"`
	// Level 地域层级
	Level int `json:"level,omitempty"`
	// Name 地域名称
	Name string `json:"name,omitempty"`
	// Parent 父级地域
	Parent uint64 `json:"parent,omitempty"`
	// Children 子级地域
	Children []uint64 `json:"children,omitempty"`
}
