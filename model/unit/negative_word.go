package unit

// AdMarketNegativeWordParam 搜索广告否词，搜索广告新增
type AdMarketNegativeWordParam struct {
	// ExactWords 精确否定词，搜索广告新增
	ExactWords []string `json:"exact_words,omitempty"`
	// PhraseWords 短语否定词，搜索广告新增
	PhraseWords []string `json:"phrase_words,omitempty"`
}
