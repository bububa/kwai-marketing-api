package creative

// ProgrammedCreativeMaterial 程序化创意物料
type ProgrammedCreativeMaterial struct {
	// Materials 物料信息
	Materials []Material `json:"materials,omitempty"`
	// StickerStyleIDs 封面贴纸类型ID
	StickerStyleIDs []int `json:"sticker_style_ids,omitempty"`
	// DescriptionTItles 广告语
	DescriptionTItles []string `json:"description_titles,omitempty"`
	// StickerTitles 封面广告语
	StickerTitles []string `json:"sticker_titles,omitempty"`
}
