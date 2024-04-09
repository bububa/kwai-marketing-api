package file

// Image 图片素材
type Image struct {
	// URL 图片预览地址
	URL string `json:"url,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// Size 图片大小
	Size int64 `json:"size,omitempty"`
	// Format 图片格式
	Format string `json:"format,omitempty"`
	// Signature 图片MD5值
	Signature string `json:"signature,omitempty"`
	// ImageToken 图片token 创建创意时使用
	ImageToken string `json:"image_token,omitempty"`
	// PicType 图片类型
	// 0-默认，5-竖版，6-横版，12-开屏，16-图集
	PicType int `json:"pic_type,omitempty"`
	// PicID 图片库图片ID
	PicID string `json:"pic_id,omitempty"`
}
