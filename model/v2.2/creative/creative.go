package creative

type Creative struct {
	CampaignID           uint64          `json:"campaign_id"`             // 广告计划 ID
	UnitID               uint64          `json:"unit_id"`                 // 广告组 ID
	CreativeID           uint64          `json:"creative_id"`             // 广告创意 ID
	CreativeName         string          `json:"creative_name"`           // 广告创意名称
	CreativeMaterialType int             `json:"creative_material_type"`  // 素材类型 0：历史创意未作区分 1：竖版视频 2：横版视频 3：后贴片单图图片创意（历史类型，已下线）4：便利贴单图图片创意 11：开屏视频 12：开屏图片
	CreativeCategory     int             `json:"creative_category"`       // 创意分类
	CreativeTag          []string        `json:"creative_tag"`            // 创意标签
	PhotoID              string          `json:"photo_id"`                // 视频作品 ID
	PhotoMD5             string          `json:"photo_md5"`               // 视频作品的md5
	MaterialURL          []string        `json:"material_url"`            // 单图创意 url
	ImageTokens          []string        `json:"image_tokens"`            // 单图创意 image_token
	Status               int             `json:"status"`                  // 广告创意状态（优先先看这个状态,计算结果） -1：不限，1：计划已暂停，3：计划超预算，6：余额不足，11：组审核中，12：组审核未通过，14：已结束，15：组已暂停，17：组超预算，19：未达投放时间，40：已删除，41：审核中，42：审核未通过，46：已暂停，52：投放中，53：作品异常，54：视频审核通过可投放滑滑场景，55：部分素材审核失败
	PutStatus            int             `json:"put_status"`              // 投放状态（操作结果） 1：投放中；2：暂停 3：删除
	ReviewDetail         string          `json:"review_detail"`           // 审核拒绝理由
	RejectVideoSnapshot  []string        `json:"reject_video_snapshot"`   // 审核拒绝图片 list 里面可以包含多个数据
	CoverURL             string          `json:"cover_url"`               // 封面 URL
	ImageToken           string          `json:"image_token"`             // 视频封面 token 若创意使用系统自动生成的首帧图片作为封面，该 token 无法复用
	CoverWidth           int64           `json:"cover_width"`             // 封面图宽度
	CoverHeight          int64           `json:"cover_height"`            // 封面图高度
	OverlayBgURL         string          `json:"overlay_bg_url"`          // 动态词包原始封面图片 URL
	OverlayBgImageToken  string          `json:"overlay_bg_image_token"`  // 动态词包原始封面图片 token
	StickerTitle         string          `json:"sticker_title"`           // 封面广告语标题
	OverlayType          string          `json:"overlay_type"`            // 贴纸样式类型
	DisplayInfo          DisplayInfo     `json:"display_info"`            //广告展示信息
	ShortSlogan          string          `json:"short_slogan"`            // 便利贴创意短广告语
	ExposeTag            string          `json:"expose_tag"`              // 广告标签
	NewExposeTag         []NewExposeTag  `json:"new_expose_tag"`          //广告标签 2 期
	ClickTrackUrl        string          `json:"click_track_url"`         // 点击监测链接 若出现与后台显示不一致，以文档为准即可
	ImpressionUrl        string          `json:"impression_url"`          // 第三方开始播放监测链接 若出现与后台显示不一致，以文档为准即可
	ActionbarClickUrl    string          `json:"actionbar_click_url"`     // 第三方点击按钮监测链接
	AdPhotoPlayedT3sUrl  string          `json:"ad_photo_played_t3s_url"` // 第三方有效播放监测链接 仅历史个别账户使用
	CreateTime           string          `json:"create_time"`             // 创建时间 格式样例："2019-06-11 15:17:25"
	UpdateTime           string          `json:"update_time"`             // 最后修改时间 格式样例："2019-06-11 15:17:25"
	PicId                string          `json:"pic_id"`                  // 图片库图片
	AppGradeType         int             `json:"app_grade_type"`          // 审核分级类型 0：默认；1：审核降级(当创意发生降级时，会限制部分流量无法投放)
	SplashPhotos         []SplashPhoto   `json:"splash_photos"`           // 开屏视频信息 creative_material_type 为 11 时
	LiveCreativeType     int             `json:"live_creative_type"`      // 粉丝直播推广创意类型 3：直投直播；4：作品引流
	SplashPictures       []SplashPicture `json:"splash_pictures"`         // 开屏图片 creative_material_type 为 12 时
	LiveTrackUrl         string          `json:"live_track_url"`          // 点击监测链接 计划 campaignType=16 粉丝直播推广时填写
	AdType               int             `json:"ad_type"`                 // 广告计划类型 0:信息流，1:搜索
	OuterLoopNative      int             `json:"outer_loop_native"`       // 是否开启原生 0关闭，1开启，不填默认为0
	KolUserType          int             `json:"kol_user_type"`           // 原生达人用户类型 2服务号原生，3聚星达人原生, 不开启原生时此项为0
	KolUserId            int64           `json:"kol_user_id"`             // 原生投放目标达人ID
	Recommendation       string          `json:"recommendation"`          // plc自定义文案
	DpaStyleTypes        []int           `json:"dpa_style_types"`         // 动态商品卡样式 14001-区域服务卡
	HighLightFlash       int             `json:"high_light_flash"`        // 高光创意状态 0：关闭 1：开启
}

type DisplayInfo struct {
	Description   string `json:"description"`     // 广告语
	ActionBarText string `json:"action_bar_text"` // 行动号召按钮文案
}

type SplashPhoto struct {
	PhotoID  string `json:"photo_id"`  // 视频ID
	PhotoMD5 string `json:"photo_md5"` // 视频的md5
	CoverURL int64  `json:"cover_url"` // 视频高度
	Height   int64  `json:"height"`    // 视频宽度
}

type SplashPicture struct {
	CoverID  int64  `json:"cover_id"`  // 封面ID
	CoverURL string `json:"cover_url"` // 封面URL
	Height   int64  `json:"height"`    // 图片高度
	Width    int64  `json:"width"`     // 图片宽度
}
