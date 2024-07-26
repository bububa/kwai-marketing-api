package creative

import "github.com/bububa/kwai-marketing-api/model"

// CreateRequest 创建创意API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativeName 创意名称; 长度为1-100字符，同一个广告组下名称不能重复
	CreativeName string `json:"creative_name,omitempty"`
	// PutStatus 广告组的投放状态; 不填时，创建的广告组为投放状态；填2时，创建的广告组为暂停状态，其他值非法
	PutStatus int `json:"put_status,omitempty"`
	// PhotoID 视频ID
	PhotoID string `json:"photo_id,omitempty"`
	// ImageToken 封面图片token; 通过上传图片接口获得，不传值则直接使用视频的首帧作为封面图片，自定义封面的图片宽高要与视频宽高一致
	ImageToken string `json:"image_token,omitempty"`
	// CreativeMaterialType 素材类型; 1：竖版视频 2：横版视频 4：便利贴单图图片创意 5： 竖版图片（优选/联盟）6：横版图片(优选/联盟/信息流/快看点) 9：小图(优选/信息流/快看点) 10：组图(优选/信息流/快看点)
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	// ImageTokens 便利贴单图图片创意token; 便利贴/图片/小图图片创意必填，目前只支持一张图片；组图图片创意需要上传3张图片，详细要求见附录
	ImageTokens []string `json:"image_tokens,omitempty"`
	// ActionBarText 行动号召按钮文案; 根据计划类型进行设置，详情见附录
	ActionBarText string `json:"action_bar_text,omitempty"`
	// Description 广告语; 长度为1-30字符，不支持换行。 如果要使用动态词包，格式如"[地区]的[男性女性]都喜欢"， 联盟广告和程序化创意不支持动态词包， 词包名可以通过下方动态词包接口获取
	Description string `json:"description,omitempty"`
	// ShortSlogan 便利贴创意短广告语; 长度为1-8字符，便利贴创意必填参数
	ShortSlogan string `json:"short_slogan,omitempty"`
	// StickerTitle 封面广告语标题; 不超过14字符， 如果要使用动态词包，格式如"[地区]的[男性女性]都喜欢"， 后贴片广告不支持动态词包， 词包名可以通过下方动态词包接口获取， sticker_title和overlay_type需同时为空值或同时非空
	StickerTitle string `json:"sticker_title,omitempty"`
	// OverlayType 贴纸样式类型; 贴纸样式可以通过查询下方贴纸掩饰接口获取， sticker_title和overlay_type需同时为空值或同时非空
	OverlayType string `json:"overlay_type,omitempty"`
	// ExposeTag 广告标签; 从 /expose_tags/list 接口获取，从下方9获取 更新的时候如果要将广告标签设置为 「暂不选择」，传空字符串即可，只支持计划类型为 2
	ExposeTag string `json:"expose_tag,omitempty"`
	// NewExposeTag 广告标签2期; 与老的expose_tag兼容, 逐渐将老的expose_tag替换掉。可以按照相关格式传递两个推荐理由 举例{“text”:""},{"text":""}
	NewExposeTag []string `json:"new_expose_tag,omitempty"`
	// SiteID 安卓下载中间页ID; 可从下方「获取已发布的下载页」接口获取（老建站6.30下线） 或通过「rest/openapi/v2/lp/page/list」获取新建站落地页（魔力建站"view_comps":7） 1.仅支持下载类广告 2.广告组选择的应用类型要为安卓 3.下载页对应的app_id要与广告组选择的app_id一致
	SiteID uint64 `json:"site_id,omitempty"`
	// ClickTrackUrl 第三方点击检测链接; 仅当广告组scene_id为1、2、6、7、10时，可选填； 广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手Android SDK时除外） 使用Marketing API创建时，监测链接使用以该文档为准
	ClickTrackUrl string `json:"click_track_url,omitempty"`
	// ImpressionUrl 第三方开始播放监测链接; 仅当广告组scene_id为3时，可选填； 广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手Android SDK时除外） 使用Marketing API创建时，监测链接使用以该文档为准
	ImpressionUrl string `json:"impression_url,omitempty"`
	// AdPhotoPlayedT3sUrl 第三方有效播放监测链接; 仅历史个别账户使用且当广告组scene_id为3时可选，与impression_url不可同时使用
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"`
	// ActionbarClickUrl 第三方点击按钮监测链接; 1.校验click_url必填的广告场景 优选(1)/信息流(2、7)/上下滑（6） 2.优化目标为激活功能必填点击监测链接,但如果安卓应用接入了快手监测sdk就不需要填写监测链接了 3.联盟场景检查click_url不能为空 4.若广告联盟的转化目标为激活，click_url、actionbar_click_url和监测SDK至少三选一
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// CreativeCategory 创意分类; 由创意分类查询接口 获得；必须是叶子结点；与创意标签同时传或同时不传 必填账户可通过接口3.17——接口获取是否必填
	CreativeCategory int `json:"creative_category,omitempty"`
	// CreativeTag 创意标签; 与创意分类参数，要么都传，要么都不传；且单个创意的创意标签最多10个；单个创意标签不能为空且不能超过10字符
	CreativeTag []string `json:"creative_tag,omitempty"`
	// LiveCreativeType 直播类型（小店直播推广类型，计划type=14）; 1 - 直投直播；2 - 视频引流直播
	LiveCreativeType int `json:"live_creative_type,omitempty"`
}

// Url implement PostRequest interface
func (r CreateRequest) Url() string {
	return "v2/creative/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
