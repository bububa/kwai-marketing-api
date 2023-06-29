package file

// Video 视频素材
type Video struct {
	// PhotoID 视频ID
	PhotoID string `json:"photo_id,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// URL 视频预览链接
	URL string `json:"url,omitempty"`
	// CoverUrl 视频首帧图片链接
	CoverUrl string `json:"cover_url,omitempty"`
	// Signature 视频md5值
	Signature string `json:"signature,omitempty"`
	// UploadTime 上传时间
	UploadTime string `json:"upload_time,omitempty"`
	// PhotoName 视频名称
	PhotoName string `json:"photo_name,omitempty"`
	// PhotoTag 视频标签
	PhotoTag string `json:"photo_tag,omitempty"`
	// NewStatus 视频状态; 0：逻辑删除，1：正常
	NewStatus int `json:"new_status,omitempty"`
	// Duration 视频时长; 单位毫秒
	Duration int64 `json:"duration,omitempty"`
	// Source 视频来源; 0：自上传，1：开眼，2：素造，7：聚星视频
	Source             int                    `json:"source,omitempty"`
	AdPhotoValuateInfo AdPhotoValuateInfoItem `json:"adPhotoValuateInfo"`
}
type AdPhotoValuateInfoItem struct {
	SimLabel                string      `json:"simLabel"`
	QualityLabel            string      `json:"qualityLabel"`
	QuotaMsg                string      `json:"quotaMsg"`
	IsDupPhoto              bool        `json:"isDupPhoto"`
	IsDelayReview           interface{} `json:"isDelayReview"`
	OptimizationSuggestions string      `json:"optimizationSuggestions"`
	RunningScore            int         `json:"runningScore"`
	HitTagCombination       int         `json:"hitTagCombination"`
}
