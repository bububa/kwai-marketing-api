package creative

// Creative 广告创意
type Creative struct {
	// CampaignID 广告计划ID
	CampaignID int64 `json:"campaign_id,omitempty"`
	// UnitID 广告组ID
	Unit int64 `json:"unit_id,omitempty"`
	// CreativeID 广告创意ID
	CreativeID int64 `json:"creative_id,omitempty"`
	// CreativeName 广告创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// CreativeMaterialType 素材类型; 0：历史创意未作区分 1：竖版视频2：横版视频3：后贴片单图图片创意（历史类型，已下线）4：便利贴单图图片创意
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	// PhotoID 视频作品ID
	PhotoID string `json:"photo_id,omitempty"`
	// MaterialUrl 单图创意url
	MaterialUrl []string `json:"material_url,omitempty"`
	// ImageTokens 单图创意image_token
	ImageTokens []string `json:"image_tokens,omitempty"`
	// Status 广告创意状态（优先先看这个状态,计算结果）;-1：不限，1：计划已暂停，3：计划超预算，6：余额不足，11：组审核中，12：组审核未通过，14：已结束，15：组已暂停，17：组超预算，19：未达投放时间，40：已删除，41：审核中，42：审核未通过，46：已暂停，52：投放中，53：作品异常，54：视频审核通过可投放滑滑场景，55：部分素材审核失败
	Status int `json:"status,omitempty"`
	// PutStatus 投放状态（操作结果）;1：投放中；2：暂停 3：删除
	PutStatus int `json:"put_status,omitempty"`
	// CreateChannel 创建渠道; 0：投放后台创建；1：Marketing API创建
	CreateChannel int `json:"create_channel,omitempty"`
	// ReviewDetail 审核拒绝理由;
	ReviewDetail string `json:"review_detail,omitempty"`
	// RejectVideoSnapshot 审核拒绝图片
	RejectVideoSnapshot []string `json:"reject_video_snapshot,omitempty"`
	// CoverUrl 封面URL
	CoverUrl string `json:"cover_url,omitempty"`
	// ImageToken 视频封面token; 若创意使用系统自动生成的首帧图片作为封面，该token无法复用
	ImageToken string `json:"image_token,omitempty"`
	// FirstFrameType 视频封面来源; 1：首帧，0：非首帧
	FirstFrameType int `json:"first_frame_type,omitempty"`
	// CoverWidth 封面图宽度
	CoverWidth int64 `json:"cover_width,omitempty"`
	// CoverHeight 封面图高度
	CoverHeight int64 `json:"cover_height,omitempty"`
	// OverlayBgUrl 动态词包原始封面图片URL
	OverlayBgUrl string `json:"overlay_bg_url,omitempty"`
	// OverlayBgImageToken 动态词包原始封面图片token
	OverlayBgImageToken string `json:"overlay_bg_image_token,omitempty"`
	// StickerTitle 封面广告语标题
	StickerTitle string `json:"sticker_title,omitempty"`
	// OverlayType 贴纸样式类型
	OverlayType string `json:"overlay_type,omitempty"`
	// DisplayInfo 广告展示信息
	DisplayInfo *DisplayInfo `json:"display_info,omitempty"`
	// ShortSlogan 便利贴创意短广告语
	ShortSlogan string `json:"short_slogan,omitempty"`
	// ExposeTag 广告标签
	ExposeTag string `json:"expose_tag,omitempty"`
	// NewExposeTag 广告标签2期
	NewExposeTag []NewExposeTag `json:"new_expose_tag,omitempty"`
	// SiteID 安卓下载中间页ID
	SiteID int64 `json:"site_id,omitempty"`
	// ClickTrackUrl 点击监测链接; 若出现与后台显示不一致，以文档为准即可
	ClickTrackUrl string `json:"click_track_url,omitempty"`
	// ImpressionUrl 第三方开始播放监测链接; 若出现与后台显示不一致，以文档为准即可
	ImpressionUrl string `json:"impression_url,omitempty"`
	// AdPhotoPlayedT3sUrl 第三方有效播放监测链接; 仅历史个别账户使用
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"`
	// ProgrammedCreativeMaterial 程序化创意物料
	ProgrammedCreativeMaterial *ProgrammedCreativeMaterial `json:"programmed_creative_material,omitempty"`
	// CreativeStatusType 否为僵尸创意; 1：僵尸 0：非僵尸
	CreativeStatusType int `json:"creative_status_type,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 最后修改时间
	UpdateTime string `json:"update_time,omitempty"`

	ActionbarClickUrl string   `json:"actionbar_click_url"`
	CreativeCategory  int      `json:"creative_category"`
	CreativeTag       []string `json:"creative_tag"`
	MerchantLibraryId int      `json:"merchant_library_id"`
	MerchantProductId string   `json:"merchant_product_id"`
	//SplashPictures             []interface{} `json:"splash_pictures"`
	//SplashPhotos               []interface{} `json:"splash_photos"`
	AdType int `json:"ad_type"`
}

type NewExposeTag struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}
