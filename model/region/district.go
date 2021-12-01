package region

// District 商圈地域定向
type District struct {
	// ID 商圈id
	ID int64 `json:"id,omitempty"`
	// Level 地域层级-商圈层级=4
	Level int `json:"level,omitempty"`
	// Name 商圈名称
	Name string `json:"name,omitempty"`
	// FullName 全名称
	FullName string `json:"full_name,omitempty"`
}
