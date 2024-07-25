package file

// Video 视频素材
type Video struct {
	// Cursor 游标查询返回游标
	Cursor int64 `json:"cursor"`
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
	PhotoTag []string `json:"photo_tag,omitempty"`
	// NewStatus 视频状态; 0：逻辑删除，1：正常
	NewStatus int `json:"new_status,omitempty"`
	// Duration 视频时长; 单位毫秒
	Duration int64 `json:"duration,omitempty"`
	// Source 视频来源; 0：自上传，1：开眼，2：素造，7：聚星视频
	Source int `json:"source,omitempty"`
	// PhotoValuateInfo 视频素材评价信息
	PhotoValuateInfo *PhotoValuateInfo `json:"photo_valuate_info,omitempty"`
	// PhotoTagIdentifyItems 视频标签内容
	PhotoTagIdentifyItems []any `json:"photo_tag_identify_items,omitempty"`
	// OuterLoopNative 是否是原生视频
	OuterLoopNative int `json:"outer_loop_native,omitempty"`
	// PhotoUserID 视频所属的userId
	PhotoUserID uint64 `json:"photo_user_id,omitempty"`
	// AtlasPicIDs 图集关联的图片ID
	AtlasPicIDs []string `json:"atlas_pic_ids,omitempty"`
	// AtlasAudioBsKey 图集关联音频的bs_key
	AtlasAudioBsKey string `json:"atlas_audio_bs_key,omitempty"`
	// AtlasAudioURL 图集音频播放链接
	AtlasAudioURL string `json:"atlas_audio_url,omitempty"`
	// ShieldStatus 视频是否已同步个人主页 1否 0是
	ShieldStatus int `json:"shield_status,omitempty"`
	// PhotoDupStatus 素材创新度
	// 0-不重复，1-重复
	PhotoDupStatus int `json:"photo_dup_status,omitempty"`
	// NativeGoodInfo 素材质量
	// 0-默认，1-良好，2-优质，3-非原生
	NativeGoodInfo int `json:"native_good_info,omitempty"`
}

// PhotoValuateInfo 视频素材评价信息
type PhotoValuateInfo struct {
	SimLabel                string      `json:"simLabel"`
	QualityLabel            string      `json:"qualityLabel"`
	QuotaMsg                string      `json:"quotaMsg"`
	IsDupPhoto              bool        `json:"isDupPhoto"`
	IsDelayReview           interface{} `json:"isDelayReview"`
	OptimizationSuggestions string      `json:"optimizationSuggestions"`
	RunningScore            int         `json:"runningScore"`
	HitTagCombination       int         `json:"hitTagCombination"`
}
