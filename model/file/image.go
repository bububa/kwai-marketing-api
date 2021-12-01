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
	// ImageToken 图片token
	ImageToken string `json:"image_token,omitempty"`
}
