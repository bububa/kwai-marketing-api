package creative

// Material 物料信息
type Material struct {
	// PhotoID 视频ID
	PhotoID string `json:"photo_id,omitempty"`
	// ImageTokens 封面图片token
	ImageTokens []string `json:"image_tokens,omitempty"`
	// CreativeMaterialType 素材类型
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	// StickerTitles 封面广告语
	StickerTitles []string `json:"sticker_titles,omitempty"`
}
