package creative

type AdvancedCreative struct {
	UnitID              int64       `json:"unit_id"`                 // 广告组ID
	PackageName         string      `json:"package_name"`            // 程序化创意包名称，1-100字符
	HorizontalPhotoIDs  []string    `json:"horizontal_photo_ids"`    // 横版视频ID列表，横版视频和竖版视频加起来只能1-5个
	VerticalPhotoIDs    []string    `json:"vertical_photo_ids"`      // 竖版视频ID列表
	CoverImageTokens    []string    `json:"cover_image_tokens"`      // 封面image_token，只能是1-4个
	CoverImageURLs      []string    `json:"cover_image_urls"`        // 封面链接地址
	StickerStyles       []int       `json:"sticker_styles"`          // 封面贴纸
	CoverSlogans        []string    `json:"cover_slogans"`           // 封面广告语
	ActionBar           string      `json:"action_bar"`              // 行动号召按钮
	Captions            []string    `json:"captions"`                // 作品广告语，只能是1-3个
	ClickURL            string      `json:"click_url"`               // 第三方点击检测链接
	ActionBarClickURL   string      `json:"actionbar_click_url"`     // 第三方ActionBar点击监控链接
	PutStatus           int         `json:"put_status"`              // 程序化创意操作状态，1：投放，2：暂停，3：删除
	ViewStatus          int         `json:"view_status"`             //程序化创意状态
	ViewStatusReason    string      `json:"view_status_reason"`      // 程序化创意状态描述
	CreateTime          string      `json:"create_time"`             // 创建时间，格式样例："2019-06-11 15:17:25"
	UpdateTime          string      `json:"update_time"`             // 更新时间，格式样例："2019-06-11 15:17:25"
	Creatives           []Creatives `json:"creatives"`               // 创建后生成的程序化创意 ID
	PicIDs              []string    `json:"pic_ids"`                 // 图片库图片ID
	AppGradeType        int         `json:"app_grade_type"`          // 审核分级类型，0：默认；1：审核降级(当创意发生降级时，会限制部分流量无法投放)
	PicList             []string    `json:"pic_list"`                // 联盟图片（横版/竖版），联盟图片imageToken
	PicURLList          []string    `json:"pic_url_list"`            // 联盟图片url（横版/竖版），联盟图片url
	PhotoList           []Photo     `json:"photo_list"`              // 素材列表
	AdPhotoPlayedT3SURL string      `json:"ad_photo_played_t3s_url"` // 第三方有效播放监测链接
	CreativeCategory    int         `json:"creative_category"`       // 创意分类
	CreativeTag         []string    `json:"creative_tag"`            // 创意标签
	NewExposeTag        []ExposeTag `json:"new_expose_tag"`          // 广告标签 2 期
}
type Photo struct {
	PhotoID              int64  `json:"photo_id"`               // 视频 ID
	CoverImageToken      string `json:"cover_image_token"`      // 封面图片 token
	CreativeMaterialType int    `json:"creative_material_type"` // 素材类型，1：竖版视频；2：横版视频
}

type ExposeTag struct {
	Text string `json:"text"`
}

type Creatives struct {
	UnitID     int64 `json:"unit_id"`     // 广告组 ID
	CreativeID int64 `json:"creative_id"` // 创意 ID
}
