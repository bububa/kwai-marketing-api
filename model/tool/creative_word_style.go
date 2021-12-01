package tool

// CreativeWordStyle 封面贴纸样式
type CreativeWordStyle struct {
	// StickerStyleID 贴纸样式ID
	StickerStyleID int `json:"sticker_style_id,omitempty"`
	// StickerStyles 贴纸样式列表
	StickerStyles []StickerStyle `json:"sticker_styles,omitempty"`
}

// StickerStyle 贴纸样式
type StickerStyle struct {
	// OverlayType 贴纸样式类型; 返回格式{stickerstyle_id}{pos}，sticker_style_id表示贴纸样式ID，pos表示位置，1: top, 2: mid, 3: bottom
	OverlayStyle string `json:"overlay_style,omitempty"`
	// OverlayPreviewUrl 贴纸样式预览链接
	OverlayPreviewUrl string `json:"overlay_preview_url,omitempty"`
}
