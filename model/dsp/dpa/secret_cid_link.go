package dpa

import "github.com/bububa/kwai-marketing-api/model"

// SecretCidLinkRequest CID服务商投放SDPA接口 API Request
type SecretCidLinkRequest struct {
	// AdvertiserID 广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// OuterID 商品ID
	OuterID string `json:"outer_id,omitempty"`
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// Description 商品描述
	Description string `json:"description,omitempty"`
	// MainImage 商品主图URL
	// 数组最大有效长度：1
	MainImage []string `json:"main_image,omitempty"`
	// ImageURLs 商品副图URLs
	ImageURLs []string `json:"image_urls,omitempty"`
	// VideoURL 商品视频
	VideoURL *VideoInfo `json:"video_url,omitempty"`
	// EnBrand 商品英文品牌名称
	EnBrand string `json:"en_brand,omitempty"`
	// Value 商品原价（元）
	// 至多精确至“厘”（小数点后三位）
	Value string `json:"value,omitempty"`
	// Saving 商品减价（元）
	// 至多精确至“厘”（小数点后三位）
	Saving string `json:"saving,omitempty"`
	// Discount 商品折扣	0.001～0.999（至多精确至小数点后三位）
	Discount string `json:"discount,omitempty"`
	// Price 商品现价（元）
	// 至多精确至“厘”（小数点后三位）
	Price string `json:"price,omitempty"`
	// SalesPromotion 商品促销活动
	SalesPromotion string `json:"sales_promotion,omitempty"`
	// DownPayment 商品首付
	DownPayment string `json:"down_payment,omitempty"`
	// Mortgage 商品月付
	Mortgage string `json:"mortgage,omitempty"`
	// DailyMortgage 商品日付
	DailyMortgage string `json:"daily_mortgage,omitempty"`
	// Address 商品销售地址
	Address string `json:"address,omitempty"`
	// Feature 商品特色卖点
	Feature string `json:"feature,omitempty"`
	// AgeV2 商品定向年龄段
	// 1表示18-23｜2表示24-30｜3表示31-40｜4表示41-49｜5表示50+
	AgeV2 []string `json:"age_v2,omitempty"`
	// Gender 商品定向性别
	// 0表示不限｜1表示男｜2表示女
	Gender int `json:"gender,omitempty"`
	// ProvinceV2 商品定向省份
	ProvinceV2 []string `json:"province_v2,omitempty"`
	// CityV2 商品定向城市
	CityV2 []string `json:"city_v2,omitempty"`
	// TargetURL 商品详情页链接
	TargetURL *UrlParamSneak `json:"target_url,omitempty"`
	// BrandURL 商品品牌页链接
	BrandURL *UrlParamSneak `json:"brand_url,omitempty"`
	// ScheduleURL 商品活动页链接
	ScheduleURL *UrlParamSneak `json:"schedule_url,omitempty"`
	// ShopKeeperURL 商品店铺页链接
	ShopKeeperURL *UrlParamSneak `json:"shop_keeper_url,omitempty"`
	// CategoryURL 商品类目页链接
	CategoryURL *UrlParamSneak `json:"category_url,omitempty"`
	// Slogan 商品广告语
	Slogan string `json:"slogan,omitempty"`
	// EcomProductInfo 其它商品信息
	EcomProductInfo *EcomProductInfo `json:"ecom_product_info,omitempty"`
	// Condition 商品新旧情况
	Condition string `json:"condition,omitempty"`
	// Comments 商品评论数
	Comments int `json:"comments,omitempty"`
	// Mark 商品评分
	// 至多精确至小数点后三位
	Mark string `json:"mark,omitempty"`
}

// EcomProductInfo 其它商品信息
type EcomProductInfo struct {
	// SpuID 标准商品单位ID
	SpuID string `json:"spu_id,omitempty"`
	// ServiceProviderName 服务商名称
	ServiceProviderName string `json:"service_provider_name,omitempty"`
	// CorporationName 商品所属营业执照
	CorporationName string `json:"corporation_name,omitempty"`
	// EcomPlatformFromService 商品来源平台
	EcomPlatformFromService string `json:"ecom_platform_from_service,omitempty"`
	// ShopNameFromService 三方平台店铺名称
	ShopNameFromService string `json:"shopF_name_from_service,omitempty"`
	// BrandNameFromService 三方平台品牌名称
	BrandNameFromService string `json:"brand_name_from_service,omitempty"`
	// FirstProductCategoryFromService 三方平台商品一级类目
	FirstProductCategoryFromService string `json:"first_product_category_from_service,omitempty"`
	// SecondProductCategoryFromService 三方平台商品二级类目
	SecondProductCategoryFromService string `json:"second_product_category_from_service,omitempty"`
	// ThirdProductCategoryFromService 三方平台商品三级类目
	ThirdProductCategoryFromService string `json:"third_product_category_from_service,omitempty"`
	// FourthProductCategoryFromService 三方平台商品四级类目
	FourthProductCategoryFromService string `json:"fourth_product_category_from_service,omitempty"`
	// ProductLabel 商品标签
	ProductLabel string `json:"product_label,omitempty"`
	// CommentTags 商品评价标签
	CommentTags string `json:"comment_tags,omitempty"`
	// ServiceTags 商品服务标签
	ServiceTags string `json:"service_tags,omitempty"`
	// AttributeTags 商品属性标签
	AttributeTags string `json:"attribute_tags,omitempty"`
	// ShopIDromService 三方平台店铺id
	ShopIDromService uint64 `json:"shop_id_from_service,omitempty"`
	// BrandIDromService 三方平台品牌id
	BrandIDromService uint64 `json:"brand_id_from_service,omitempty"`
	// FirstProductCategoryIDFromService 三方平台商品一级类目id
	FirstProductCategoryIDFromService uint64 `json:"first_product_category_id_from_service,omitempty"`
	// SecondProductCategoryIDFromService 三方平台商品二级类目id
	SecondProductCategoryIDFromService uint64 `json:"second_product_category_id_from_service,omitempty"`
	// ThirdProductCategoryIDFromService 三方平台商品三级类目id
	ThirdProductCategoryIDFromService uint64 `json:"third_product_category_id_from_service,omitempty"`
	// FourthProductCategoryIDFromService 三方平台商品四级类目id
	FourthProductCategoryIDFromService uint64 `json:"fourth_product_category_id_from_service,omitempty"`
	// Bought 商品销售量
	Bought int `json:"bought,omitempty"`
}

// Url implement PostRequest interface
func (r SecretCidLinkRequest) Url() string {
	return "gw/dsp/dpa/secret/cid/link"
}

// Encode implement PostRequest interface
func (r SecretCidLinkRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
