package creative

// AdvancedCreative 程序化创意
type AdvancedCreative struct {
	// UnitID 广告组 ID
	// 一个组下只能有一个程序化创意，只有这个广告组的 unit_type 为 7 才能创建程序化创意
	UnitID uint64 `json:"unit_id,omitempty"`
	// PackageName 程序化创意名称
	PackageName string `json:"package_name,omitempty"`
	// HorizontalPhotoIDs 横版视频 id list
	HorizontalPhotoIDs []string `json:"horizontal_photo_ids,omitempty"`
	// VerticalPhotoIDs 竖版视频 id list
	VerticalPhotoIDs []string `json:"vertical_photo_ids,omitempty"`
	// CoverImageTokens 封面 image_token
	CoverImageTokens []string `json:"cover_image_tokens,omitempty"`
	// CoverImageURLs 封面链接地址
	CoverImageURLs []string `json:"cover_image_urls,omitempty"`
	// StickerStyles 贴纸样式
	StickerStyles []int `json:"sticker_styles,omitempty"`
	// CoverSlogans 封面 slogans
	CoverSlogans []string `json:"cover_slogans,omitempty"`
	// ActionBar 行动号召按钮
	ActionBar string `json:"action_bar,omitempty"`
	// Captions 作品广告语
	// 每个不超过 30 个字符，英文字符两个算一个字符，最多可传 3 个
	Captions []string `json:"captions,omitempty"`
	// ClickURL 第三方点击检测链接
	// 不能超过 1024 字符 ocpx_action_type 是 180 并且应用没有接入 sdk，监测链接必填； 计划 type 是 2（推广应用安装），ocpx_action_type 是注册（396）、付费（190）、完件（384）、授信（383），并且没有接入 sdk，监测链接必填
	ClickURL string `json:"click_url,omitempty"`
	// ActionbarClickURL 第三方 ActionBar 点击监控链接，命中有效触点白名单表示有效触点监测链接（限：快手主站/极速版场景）
	// 部分客户使用 actionbar_click_url 不为空时，click_url 必填，不能超过 1024 字符
	ActionbarClickURL string `json:"actionbar_click_url,omitempty"`
	// PutStatus 程序化创意操作状态
	// 1：投放，2：暂停，3：删除
	PutStatus int `json:"put_status,omitempty"`
	// ViewStatus 程序化创意状态
	// -1：不限，1：计划已暂停，3：计划超预算，5：计划已删除，6：余额不足，11：组审核中，12：组审核未通过，14：已结束，15：组已暂停，17：组超预算，19：未达投放时间，22：不在投放时段，40：创意已删除，41：审核中，42：审核未通过，46：已暂停，52：投放中，53：作品异常，55：部分素材审核通过，56：部分审核失败，62：待送审
	ViewStatus int `json:"view_status,omitempty"`
	// ViewStatusReason 程序化创意状态描述
	ViewStatusReason string `json:"view_status_reason,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	// Creatives 创建后生成的程序化创意 ID
	Creatives []Creative `json:"creatives,omitempty"`
	// PicIDs 图片库图片ID
	PicIDs []string `json:"pic_ids,omitempty"`
	// AppGradeType 审核分级类型
	// 0：默认；1：审核降级(当创意发生降级时，会限制部分流量无法投放)
	AppGradeType int `json:"app_grade_type,omitempty"`
	// PicList 联盟图片（横版/竖版）
	// 联盟图片imageToken
	PicList []string `json:"pic_list,omitempty"`
	// PicURLList 联盟图片url（横版/竖版）
	PicURLList []string `json:"pic_url_list,omitempty"`
	// PhotoList 素材列表
	PhotoList []Photo `json:"photo_list,omitempty"`
	// AdPhotoPlayedT3sURL 第三方有效播放监测链接
	// 白名单功能，且当广告组 scene_id 为 27（开屏） 时不支持该检测链接；与 impression_url 不可同时使用
	AdPhotoPlayedT3sURL string `json:"ad_photo_played_t3s_url,omitempty"`
	// CreativeCategory 创意分类
	// 由创意分类查询接口 获得；必须是叶子结点；与创意标签同时传或同时不传 可通过工具-功能名单-获取创意分类标签白名单客户接口获取是否必填。注：不可传入负值
	CreativeCategory int `json:"creative_category,omitempty"`
	// CreativeTag 创意标签
	// 与创意分类参数，要么都传，要么都不传；且单个创意的创意标签最多 10 个；单个创意标签不能为空且不能超过 10 字符，
	CreativeTag []string `json:"creative_tag,omitempty"`
	// NewExposeTag 广告标签 2 期
	// 按照相关格式传递两个推荐理由 举例{“text”:"工厂直发"},{"text":"限时专享"}
	NewExposeTag []NewExposeTag `json:"new_expose_tag,omitempty"`
	// MaterialIntelligentOptimize 素材智能优化开关
	// 0-关闭，1-开启，不传默认关闭。仅白名单用户可以使用。
	MaterialIntelligentOptimize int `json:"material_intelligent_optimize,omitempty"`
	// OuterLoopNative 是否开启原生
	// 1开启，0关闭，不填则默认为0，投放快手信息流广告时（ad_type=0或默认不填，scene_id包含1优选广告位、6上下滑大屏广告、7双列信息流广告），当campaignType=2提升应用安装、5收集销售线索、7提升应用活跃、19推广快手小程序、30快手号-短剧推广时，可开启原生投放。注：投放快手信息流广告且升级白名单账户必选 outer_loop_native = 1，否则会报错
	OuterLoopNative int `json:"outer_loop_native,omitempty"`
	// KOLUserType 原生达人用户类型
	// 1普通快手号原生，2服务号原生，3聚星达人原生，当outer_loop_native=1时此项必填,当 campaignType=30只能填写1普通快手号
	KOLUserType int `json:"kol_user_type,omitempty"`
	// KOLUserID 原生投放目标达人ID
	// 开启原生场景下必传，即当outer_loop_native=1时此项必填， 计划 campaignType=30 短剧推广时，值为短剧作者ID
	KOLUserID uint64 `json:"kol_user_id,omitempty"`
	// Recommendation plc自定义文案
	// 开启原生时可用
	Recommendation string `json:"recommendation,omitempty"`
}
