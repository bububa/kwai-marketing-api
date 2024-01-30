package creative

import "github.com/bububa/kwai-marketing-api/model"

// AdvancedCreativeCreateRequest 创建程序化创意 API Request
type AdvancedCreativeCreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组 ID
	// 一个组下只能有一个程序化创意，只有这个广告组的 unit_type 为 7 才能创建程序化创意
	UnitID uint64 `json:"unit_id,omitempty"`
	// PackageName 程序化创意名称
	// 不能为空，1-100 字符
	PackageName string `json:"package_name,omitempty"`
	// StickerStyles 封面贴纸
	// （仅搜索广告支持）如果使用封面贴纸 sticker_Styles 和 cover_slogans 必须同时传值，最多选择 6 个
	StickerStyles []int `json:"sticker_styles,omitempty"`
	// CoverSlogans 封面口号
	// （仅搜索广告支持） 0-14 字符，最多选择 6 个（每个中文和英文字符都算一个字符）
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
	// AdPhotoPlayedT3sURL 第三方有效播放监测链接
	// 白名单功能，且当广告组 scene_id 为 27（开屏） 时不支持该检测链接；与 impression_url 不可同时使用
	AdPhotoPlayedT3sURL string `json:"ad_photo_played_t3s_url,omitempty"`
	// CreativeCategory 创意分类
	// 由创意分类查询接口 获得；必须是叶子结点；与创意标签同时传或同时不传 可通过工具-功能名单-获取创意分类标签白名单客户接口获取是否必填。注：不可传入负值
	CreativeCategory int `json:"creative_category,omitempty"`
	// CreativeTag 创意标签
	// 与创意分类参数，要么都传，要么都不传；且单个创意的创意标签最多 10 个；单个创意标签不能为空且不能超过 10 字符，
	CreativeTag []string `json:"creative_tag,omitempty"`
	// PhotoList 素材列表
	// 新创建程序化创意请使用此参数，最多支持 10 组素材(传递后将忽略 horizontal_photo_ids,vertical_photo_ids,cover_image_tokens，7.15 日后老字段下线)
	PhotoList []Photo `json:"photo_list,omitempty"`
	// PicList 联盟图片（横版/竖版）
	// 需要传入image_token列表，image_token通过上传图片接口获取
	PicList []string `json:"pic_list,omitempty"`
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

// Url implement PostRequest interface
func (r AdvancedCreativeCreateRequest) Url() string {
	return "gw/dsp/advanced_creative/create"
}

// Encode implement PostRequest interface
func (r AdvancedCreativeCreateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdvancedCreativeCreateResponse 创建程序化创意 API Response
type AdvancedCreativeCreateResponse struct {
	// UnitID 广告组 id
	UnitID uint64 `json:"unit_id,omitempty"`
}
