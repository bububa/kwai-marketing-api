package creative

import "encoding/json"

// AdvancedProgramUpdateRequest 修改程序化2.0创意
type AdvancedProgramUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:'advertiser_id,omitempty'`
	// UnitID 广告组ID
	UnitID int64 `json:"unit_id,omitempty"`
	// PackageName 程序化创意包名称，1-100 字符，
	PackageName string `json:"package_name,omitempty"`
	// HorizontalPhotoIDs 横版视频 id list; 横版视频和竖版视频加起来只能 1-5 个
	HorizontalPhotoIDs []string `json:"horizontal_photo_ids,omitempty"`
	// VerticalPhotoIDs 竖版视频 id list
	VerticalPhotoIDs []string `json:"vertical_photo_ids,omitempty"`
	// CoverImageTokens 封面 image_token;只能是 1-4 个
	CoverImageTokens []string `json:"cover_image_tokens,omitempty"`
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
	// CreativeCategory 创意分类; 由创意分类查询接口获得； 必须是叶子结点；与创意标签同时传或同时不传
	CreativeCategory int `json:"creative_category,omitempty"`
	// PhotoList 素材列表; 新创建程序化创意请使用此参数，最多支持10组素材(传递后将忽略horizontal_photo_ids,vertical_photo_ids,cover_image_tokens，7.15日后老字段下线)
	PhotoList []Photo `json:"photo_list,omitempty"`
}

// Url implement PostRequest interface
func (r AdvancedProgramUpdateRequest) Url() string {
	return "creative/advanced/program/update"
}

// Encode implement PostRequest interface
func (r AdvancedProgramUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
