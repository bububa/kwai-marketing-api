package tool

// CreativeWord 动态词包
type CreativeWord struct {
	// Name 创意词包名称
	Name string `json:"name,omitempty"`
	// ReplaceWords 替换词
	ReplaceWords []string `json:"replace_words,omitempty"`
	// DefaultWord 默认词
	DefaultWord string `json:"default_word,omitempty"`
	// MaxWordLength 替换词最大长度
	MaxWordLength int `json:"max_word_length,omitempty"`
}
