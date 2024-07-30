package creative

// Creative 自定义创意
type Creative struct {
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativeID 创意 ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeName 创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// CreativeMaterialType 素材类型
	// 1：竖版视频；2：横版视频；5： 竖版图片（优选/联盟）；6：横版图片(优选/联盟/信息流/快看点)；9：小图(优选/信息流/快看点)；10：组图(优选/信息流/快看点) ；11:开屏视频；12：开屏图片；14：DPA模板（联盟）。搜索广告当前仅支持1、2。直播间直投创意(live_creative_type = 3 或者 (campaign_type = 14 && live_creative_type = 1))时可不传。
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	// CreativeCategory 创意分类
	// 由创意分类查询接口 获得；必须是叶子结点；与创意标签同时传或同时不传 可通过工具-功能名单-获取创意分类标签白名单客户接口获取是否必填。注：不可传入负值
	CreativeCategory int `json:"creative_category,omitempty"`
	// CreativeTag 创意标签
	// 与创意分类参数，要么都传，要么都不传；且单个创意的创意标签最多 10 个；单个创意标签不能为空且不能超过 10 字符，
	CreativeTag []string `json:"creative_tag,omitempty"`
	// PhotoID 视频 ID
	PhotoID string `json:"photo_id,omitempty"`
	// PhtoMD5 视频作品的md5
	PhotoMD5 string `json:"photo_md5,omitempty"`
	// MaterialURL 单图创意 url
	MaterialURL []string `json:"material_url,omitempty"`
	// ImageTokens 单图创意 image_token
	ImageTokens []string `json:"image_tokens,omitempty"`
	// Status 广告创意状态（优先先看这个状态,计算结果）
	// -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，11：组审核中，12：组审核未通过，14：已结束，15：组已暂停，17：组超预算，19：未达投放时间，40：已删除，41：审核中，42：审核未通过，46：已暂停，52：投放中，53：作品异常，54：视频审核通过可投放滑滑场景，55：部分素材审核失败
	Status int `json:"status,omitempty"`
	// PutStatus 投放状态（操作结果）
	// 	1：投放中；2：暂停 3：删除
	PutStatus int `json:"put_status,omitempty"`
	// ReviewDetail 审核拒绝理由
	ReviewDetail string `json:"review_detail,omitempty"`
	// ReviewVideoSnapshot 审核拒绝图片
	// list 里面可以包含多个数据
	ReviewVideoSnapshot []string `json:"review_video_snapshot,omitempty"`
	// CoverURL 封面 URL
	CoverURL string `json:"cover_url,omitempty"`
	// ImageToken 视频封面 token
	// 若创意使用系统自动生成的首帧图片作为封面，该 token 无法复用
	ImageToken string `json:"image_token,omitempty"`
	// CoverWidth 封面图宽度
	CoverWidth int64 `json:"cover_width,omitempty"`
	// CoverHeight 封面图高度
	CoverHeight int64 `json:"cover_height,omitempty"`
	// OverlayBgURL 动态词包原始封面图片 URL
	OverlayBgURL string `json:"overlay_bg_url,omitempty"`
	// OverlayBgImageToken 动态词包原始封面图片 token
	OverlayBgImageToken string `json:"overlay_bg_image_token,omitempty"`
	// StickerTitle 封面广告语标题
	StickerTitle string `json:"sticker_title,omitempty"`
	// OverlayType 贴纸样式类型
	OverlayType string `json:"overlay_type,omitempty"`
	// DisplayInfo 广告展示信息
	DisplayInfo *DisplayInfo `json:"display_info,omitempty"`
	// ShortSlogen 便利贴创意短广告语
	ShortSlogen string `json:"short_slogen,omitempty"`
	// ExposeTag 广告标签
	ExposeTag string `json:"expose_tag,omitempty"`
	// NewExposeTag 广告标签 2 期
	NewExposeTag []NewExposeTag `json:"new_expose_tag,omitempty"`
	// ClickTrackURL 点击监测链接
	ClickTrackURL string `json:"click_track_url,omitempty"`
	// ImpressionURL 第三方开始播放监测链接
	// 广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手 Android SDK 时除外） 使用 Marketing API 创建时，监测链接使用以该文档为准
	ImpressionURL string `json:"impression_url,omitempty"`
	// AdPhotoPlayedT3sURL 第三方有效播放监测链接
	// 白名单功能，且当广告组 scene_id 为 27（开屏） 时不支持该检测链接；与 impression_url 不可同时使用
	AdPhotoPlayedT3sURL string `json:"ad_photo_played_t3s_url,omitempty"`
	// ActionbarClickURL 第三方点击按钮监测链接，命中有效触点白名单表示有效触点监测链接（限：快手主站/极速版场景）
	// 1.校验 click_url 必填的广告场景 优选(1)/信息流(2、7)/上下滑（6） 2.优化目标为激活功能必填点击监测链接,但如果安卓应用接入了快手监测 sdk 就不需要填写监测链接了 3.联盟场景检查 click_url 不能为空 4.若广告联盟的转化目标为激活，click_url、actionbar_click_url 和监测 SDK 至少三选一
	ActionbarClickURL string `json:"actionbar_click_url,omitempty"`
	// CreateTime 创建时间
	// 格式样例："2019-06-11 15:17:25"
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 最后修改时间
	UpdateTime string `json:"update_time,omitempty"`
	// PicID 图片库图片
	PicID string `json:"pic_id,omitempty"`
	// AppGradeType 审核分级类型
	// 0：默认；1：审核降级(当创意发生降级时，会限制部分流量无法投放)
	AppGradeType int `json:"app_grade_type,omitempty"`
	// SplashPhotos 开屏视频信息
	// creative_material_type 为 11 时
	SplashPhotos []SplashPhoto `json:"splash_photos,omitempty"`
	// LiveCreativeType 粉丝直播推广创意类型
	// 3：直投直播；4：作品引流
	LiveCreativeType int `json:"live_creative_type,omitempty"`
	// SplashPictures 开屏图片
	// creative_material_type 为 12 时
	SplashPictures []SplashPicture `json:"splash_pictures,omitempty"`
	// LiveTrackURL 点击监测链接
	// 计划 campaignType=16 粉丝直播推广时填写
	LiveTrackURL string `json:"live_track_url,omitempty"`
	// AdType 广告计划类型
	// 0:信息流，1:搜索
	AdType int `json:"ad_type,omitempty"`
	// OuterLoopNative 是否开启原生
	// 0关闭，1开启，不填默认为0
	OuterLoopNative int `json:"outer_loop_native,omitempty"`
	// KOLUserType 原生达人用户类型
	// 1普通快手号原生，2服务号原生，3聚星达人原生，当outer_loop_native=1时此项必填,当 campaignType=30只能填写1普通快手号
	KOLUserType int `json:"kol_user_type,omitempty"`
	// KOLUserID 原生投放目标达人ID
	// 开启原生场景下必传，即当outer_loop_native=1时此项必填， 计划 campaignType=30 短剧推广时，值为短剧作者ID
	KOLUserID uint64 `json:"kol_user_id,omitempty"`
	// Recommendation plc自定义文案
	Recommendation string `json:"recommendation,omitempty"`
	// DpaStyleTypes 动态商品卡样式
	DpaStyleTypes []int `json:"dpa_style_types,omitempty"`
	// HigLightFlash 高光创意状态
	// 0：关闭 1：开启
	HighLightFlash int `json:"high_light_flash,omitempty"`
	// MaterialIntelligentOptimize 素材智能优化开关
	// 0-关闭，1-开启，不传默认关闭。仅白名单用户可以使用。
	MaterialIntelligentOptimize int `json:"material_intelligent_optimize,omitempty"`
	// CreativeMode 创意生成模式
	// 0：默认-普通模式 1：AIGC生成
	CreativeMode int `json:"creative_mode,omitempty"`
	// OpenAccountNative 是否为原生扩量
	// 0：否 1：是
	OpenAccountNative int `json:"open_account_native,omitempty"`
}

// DisplayInfo 广告展示信息
type DisplayInfo struct {
	// Description 广告语
	Description string `json:"description,omitempty"`
	// ActionBarText 行动号召按钮文案
	ActionBarText string `json:"action_bar_text,omitempty"`
}

// NewExposeTag 广告标签 2 期
type NewExposeTag struct {
	// Text 广告标签text
	Text string `json:"text,omitempty"`
	// URL 广告标签url
	URL string `json:"url,omitempty"`
}

// Photo 素材
type Photo struct {
	// PhotoID 视频 ID
	PhotoID uint64 `json:"photo_id,omitempty"`
	// CoverImageToken 封面图片 token
	// 通过上传图片接口获得，不传值则直接使用视频的首帧作为封面图片，自定义封面的图片宽高要与视频宽高一致，使用智能抽帧时不需要传递。
	CoverImageToken string `json:"cover_image_token,omitempty"`
	// CreativeMaterialType 素材类型
	// 1：竖版视频 2：横版视频
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
}

// SplashPhoto 开屏视频信息
type SplashPhoto struct {
	// PhotoID 视频ID
	PhotoID string `json:"photo_id,omitempty"`
	// PhotoMD5 视频的md5
	PhotoMD5 string `json:"photo_md5,omitempty"`
	// Width 视频高度
	Width int64 `json:"width,omitempty"`
	// Height 视频宽度
	Height int64 `json:"height,omitempty"`
}

// SplashPicture 开屏图片
type SplashPicture struct {
	// CoverID 封面ID
	CoverID uint64 `json:"cover_id,omitempty"`
	// CoverURL 封面URL
	CoverURL string `json:"cover_url,omitempty"`
	// Height 图片高度
	Height int64 `json:"height,omitempty"`
	// Width 图片宽度
	Width int64 `json:"width,omitempty"`
}
