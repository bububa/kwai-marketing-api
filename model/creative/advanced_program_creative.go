package creative

// AdvancedProgramCreative 获取程序化创意2.0信息
type AdvancedProgramCreative struct {
	// UnitID 广告组ID
	Unit int64 `json:"unit_id,omitempty"`
	// PackageName 程序化创意包名称，1-100 字符，
	PackageName string `json:"package_name,omitempty"`
	// HorizontalPhotoIDs 横版视频 id list; 横版视频和竖版视频加起来只能 1-5 个
	HorizontalPhotoIDs []string `json:"horizontal_photo_ids,omitempty"`
	// VerticalPhotoIDs 竖版视频 id list
	VerticalPhotoIDs []string `json:"vertical_photo_ids,omitempty"`
	// CoverImageTokens 封面 image_token;只能是 1-4 个
	CoverImageTokens []string `json:"cover_image_tokens,omitempty"`
	// CoverImageUrls 封面链接地址
	CoverImageUrls []string `json:"cover_image_urls,omitempty"`
	// SiteID 建站 id
	SiteID int64 `json:"site_id,omitempty"`
	// StickerStyles 封面贴纸
	StickerStyles []int `json:"sticker_styles,omitempty"`
	// CoverSlogans 封面广告语
	CoverSlogans []string `json:"cover_slogans,omitempty"`
	// ActionBar 行动号召按钮
	ActionBar string `json:"action_bar,omitempty"`
	// Captions 作品广告语; 只能是 1-3 个
	Captions []string `json:"captions,omitempty"`
	// ClickUrl 第三方点击检测链接
	ClickUrl string `json:"click_url,omitempty"`
	// ActionbarClickUrl 第三方ActionBar点击监控链接
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// PutStatus 程序化创意操作状态，1：投放，2：暂停，3：删除
	PutStatus int `json:"put_status,omitempty"`
	// ViewStatus 程序化创意状态; -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，11：组审核中，12：组审核未通过，14：已结束，15：组已暂停，17：组超预算，19：未达投放时间，40：创意已删除，41：审核中，42：审核未通过，46：已暂停，52：投放中，53：作品异常，55：部分素材审核通过
	ViewStatus int `json:"view_status,omitempty"`
	// ViewStatusReason 程序化创意状态描述
	ViewStatusReason string `json:"view_status_reason,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 最后修改时间
	UpdateTime string `json:"update_time,omitempty"`
}
