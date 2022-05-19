package target

// Keyword 行为定向关键词
type Keyword struct {
	// ID 关键词id
	ID uint64 `json:"id,omitempty"`
	// Name 关键词名称
	Name string `json:"name,omitempty"`
}
