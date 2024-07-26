package creative

import "encoding/json"

// CreateRequest 创建创意API Request
type CreateRequest struct {
	//advertiser_id	long	必填	广告主 ID	在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//unit_id	long	必填	广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	//creative_name	string	必填	创意名称	长度为 1-100 字符，同一个广告组下名称不能重复
	CreativeName string `json:"creative_name,omitempty"`
	//photo_id	string	选填	视频 ID
	PhotoID string `json:"photo_id,omitempty"`
	//image_token	string	选填	封面图片 token	通过上传图片接口获得，不传值则直接使用视频的首帧作为封面图片，自定义封面的图片宽高要与视频宽高一致
	ImageToken string `json:"image_token,omitempty"`
	//dpa_template_id	long	选填	DPA模板ID	creative_material_type = 14 时，必填。通过DPA 模板信息接口取
	DPATemplateID uint64 `json:"dpa_template_id,omitempty"`
	//creative_material_type	int	必填	素材类型	1：竖版视频；2：横版视频；4：便利贴单图图片创意；5： 竖版图片（优选/联盟）；6：横版图片(优选/联盟/信息流/快看点)；9：小图(优选/信息流/快看点)；10：组图(优选/信息流/快看点) ；11:开屏视频；12：开屏图片；14：DPA模板（联盟）。搜索广告当前仅支持1、2
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	//image_tokens	string[]	选填	便利贴单图图片创意 token	便利贴/图片/小图图片创意必填，目前只支持一张图片；组图图片创意需要上传 3 张图片，详细要求见附录
	ImageTokens []string `json:"image_tokens,omitempty"`
	//action_bar_text	string	必填	行动号召按钮文案	根据计划类型进行设置，详情见附录
	ActionBarText string `json:"action_bar_text,omitempty"`
	//description	string	必填	广告语	长度为 1-30 字符，不支持换行。 如果要使用动态词包，格式如"[地区]的[男性女性]都喜欢"， 联盟广告和程序化创意不支持动态词包， 词包名可以通过下方动态词包接口获取
	Description string `json:"description,omitempty"`
	//new_expose_tag	struct[]	选填	广告标签 2 期	按照相关格式传递两个推荐理由 举例{“text”:"工厂直发"},{"text":"限时专享"}
	NewExposeTag []NewExposeTag `json:"new_expose_tag,omitempty"`
	//click_track_url	string	选填	第三方点击检测链接	仅当广告组 scene_id 为 1、2、6、7、10 时，可选填； 广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手 Android SDK 时除外） 使用 Marketing API 创建时，监测链接使用以该文档为准
	ClickTrackURL string `json:"click_track_url,omitempty"`
	//impression_url	string	选填	第三方开始播放监测链接	广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手 Android SDK 时除外） 使用 Marketing API 创建时，监测链接使用以该文档为准
	ImpressionURL string `json:"impression_url,omitempty"`
	//ad_photo_played_t3s_url	string	选填	第三方有效播放监测链接
	AdPhotoPlayedT3SURL string `json:"ad_photo_played_t3s_url,omitempty"`
	//actionbar_click_url	string	选填	第三方点击按钮监测链接	1.校验 click_url 必填的广告场景 优选(1)/信息流(2、7)/上下滑（6） 2.优化目标为激活功能必填点击监测链接,但如果安卓应用接入了快手监测 sdk 就不需要填写监测链接了 3.联盟场景检查 click_url 不能为空 4.若广告联盟的转化目标为激活，click_url、actionbar_click_url 和监测 SDK 至少三选一
	ActionbarClickURL string `json:"actionbar_click_url,omitempty"`
	//creative_category	int	金融，教育，游戏，小说，电商 如上行业必填	创意分类	由创意分类查询接口 获得；必须是叶子结点；与创意标签同时传或同时不传 可通过工具-功能名单-获取创意分类标签白名单客户接口获取是否必填。注：不可传入负值
	CreativeCategory int `json:"creative_category,omitempty"`
	//creative_tag	string[]	选创意分类 必填	创意标签	与创意分类参数，要么都传，要么都不传；且单个创意的创意标签最多 10 个；单个创意标签不能为空且不能超过 10 字符，
	CreativeTag []string `json:"creative_tag,omitempty"`
	//live_creative_type	int	选填	直播类型（粉丝直播推广计划type=16时必填）	3 - 直投直播；4 - 视频引流直播
	LiveCreativeType int `json:"live_creative_type,omitempty"`
	//live_track_url	string	选填	点击监测链接	计划 campaignType=16 粉丝直播推广时可填写
	LiveTrackURL string `json:"live_track_url,omitempty"`
	//dpa_style_types	int[]	选填	动态商品卡样式	14001-区域服务卡
	DpaStyleTypes []int `json:"dpa_style_types,omitempty"`
	//outer_loop_native	int	选填	是否开启原生	0关闭，1开启，不填默认为0仅在计划 campaignType=2提升应用安装、5收集销售线索、7提升应用活跃、19推广快手小程序时，可开启原生投放。
	OuterLoopNative int `json:"outer_loop_native,omitempty"`
	//kol_user_type	int	选填	原生达人用户类型	2服务号原生，3聚星达人原生，当outer_loop_native为1时此项必填
	KolUserType int `json:"kol_user_type,omitempty"`
	//kol_user_id 原生投放目标达人ID	 开启原生场景下必传，即当outer_loop_native=1时此项必填
	KolUserId int64 `json:"kol_user_id,omitempty"`
	//high_light_flash	int	选填	高光创意状态	0：关闭 1：开启
	HighLightFlash int `json:"high_light_flash,omitempty"`
	//splash_photo_ids	sting[]	选填	开屏视频 id	creative_material_type 为 11 时必填，使用上传视频接口时返回的 photo_id；即素材类型是开屏视频时，必须传入全尺寸的 4 条视频，具体参考素材层级接口
	SplashPhotoIDs []string `json:"splash_photo_ids,omitempty"`
	//splash_image_tokens	sting[]	选填	开屏图片 token	creative_material_type 为 12 时必填，使用上传图片接口时返回的 image_token，素材类型是开屏图片时，必须传入全尺寸的 6 张图片，具体参考素材层级接口
	SplashImageTokens []string `json:"splash_image_tokens,omitempty"`
}

// NewExposeTag 广告标签
type NewExposeTag struct {
	Text string `json:"text,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Url implement PostRequest interface
func (r CreateRequest) Url() string {
	return "gw/dsp/creative/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
