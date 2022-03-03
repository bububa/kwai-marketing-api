package creative

import "encoding/json"

// CreateAdvancedRequest 创建程序化2.0创意 API Request
type CreateAdvancedRequest struct {
	// AdvertiserId 广告主ID
	AdvertiserId int64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组ID
	Unit int64 `json:"unit_id,omitempty"`
	//PackageName 程序化创意名称
	PackageName string `json:"package_name"`
	// SiteID 安卓下载中间页ID
	SiteID int64 `json:"site_id,omitempty"`
	// StickerStyles 封面贴纸
	StickerStyles []int `json:"sticker_styles,omitempty"`
	// CoverSlogans 封面广告语
	CoverSlogans []string `json:"cover_slogans,omitempty"`
	// ActionBar 行动号召按钮
	ActionBar string `json:"action_bar"`
	// Captions 作品广告语
	Captions []string `json:"captions"`
	// ClickUrl 点击监测链接; 若出现与后台显示不一致，以文档为准即可
	ClickUrl string `json:"click_url,omitempty"`
	// ActionbarClickUrl 第三方 ActionBar 点击监控链接
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// CreativeCategory 创意分类
	CreativeCategory int `json:"creative_category,omitempty"`
	// CreativeTag 创意标签
	CreativeTag []string `json:"creative_tag,omitempty"`
	// PhotoList 素材列表
	PhotoList []PhotoItem `json:"photo_list,omitempty"`
	// PicList 联盟图片（横版/竖版）
	PicList []string `json:"pic_list,omitempty"`
}

type PhotoItem struct {
	// CreativeMaterialType 素材类型; 0：历史创意未作区分 1：竖版视频2：横版视频3：后贴片单图图片创意（历史类型，已下线）4：便利贴单图图片创意
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
	// PhotoID 视频作品ID
	PhotoID string `json:"photo_id,omitempty"`
	// CoverImageTokens 封面图片cover_image_token
	CoverImageTokens []string `json:"cover_image_tokens,omitempty"`
}

// Url implement PostRequest interface
func (r CreateAdvancedRequest) Url() string {
	return "v2/creative/advanced/program/create"
}

// Encode implement PostRequest interface
func (r CreateAdvancedRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
