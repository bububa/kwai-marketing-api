package dpa

// ProductInfo 商品信息
type ProductInfo struct {
	// IndustryID 商品所属行业ID 1-电商, 3-保险, 5-书库, 8-音视频, 9-医药, 10-招聘, 11-区域服务, 12-通信
	IndustryID int `json:"industry_id,omitempty"`
	// SubIndustryID 商品类型ID(所属二级行业ID)
	SubIndustryID string `json:"sub_industry_id,omitempty"`
	// IndustryName 商品行业名称 1-电商, 3-保险, 5-书库, 8-音视频, 9-医药, 10-招聘, 11-区域服务, 12-通信
	IndustryName string `json:"industry_name,omitempty"`
	// SubIndustryName 商品类型名称(所属二级行业名称) 小说: 1-网络小说, 2-漫画, 3-有声读物, 4-网赚; 音视频: 电影, 动漫, 纪录片, 综艺, 电视剧, 短剧, 音乐, 直播, 网赚, 其他
	SubIndustryName string `json:"sub_industry_name,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// OuterID 商品第三方ID
	OuterID string `json:"outer_id,omitempty"`
	// ProductID 商品ID 快手生成商品唯一标识
	ProductID uint64 `json:"product_id,omitempty"`
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// Description 商品描述
	Description string `json:"description,omitempty"`
	// MainImageURL 商品主图URL
	MainImageURL []string `json:"main_image_url,omitempty"`
	// ImageURLs 商品副图URLs
	ImageURLs []string `json:"image_urls,omitempty"`
	// Video 商品视频信息
	Video VideoInfo `json:"video,omitempty"`
	// TargetURLs 商品链接信息
	TargetURLs UrlParamSneak `json:"target_urls,omitempty"`
	// Condition 商品新旧情况 全新, 翻新, 二手
	Condition string `json:"condition,omitempty"`
	// Label 商品标签
	Label []string `json:"label,omitempty"`
	// FirstCategoryID 商品一级类目ID
	FirstCategoryID uint64 `json:"first_category_id,omitempty"`
	// FirstCategoryName 商品一级类目名称
	FirstCategoryName string `json:"first_category_name,omitempty"`
	// SecondCategoryID 商品二级类目ID
	SecondCategoryID uint64 `json:"second_category_id,omitempty"`
	// SecondCategoryName 商品二级类目名称
	SecondCategoryName string `json:"second_category_name,omitempty"`
	// ThirdCategoryID 商品三级类目ID
	ThirdCategoryID uint64 `json:"third_category_id,omitempty"`
	// ThirdCategoryName 商品三级类目名称
	ThirdCategoryName string `json:"third_category_name,omitempty"`
	// CategoryID 商品叶子类目ID
	CategoryID uint64 `json:"category_id,omitempty"`
	// CategoryName 商品叶子类目名称
	CategoryName string `json:"category_name,omitempty"`
	// BrandID 商品品牌ID
	BrandID string `json:"brand_id,omitempty"`
	// BrandName 商品品牌名
	BrandName string `json:"brand_name,omitempty"`
	// EnBrandName 商品英文品牌名
	EnBrandName string `json:"en_brand_name,omitempty"`
	// Value 商品原价 最小精度小数点后3位
	Value float64 `json:"value,omitempty"`
	// Price 商品现价 最小精度小数点后3位
	Price float64 `json:"price,omitempty"`
	// Saving 商品优惠价格 最小精度小数点后3位
	Saving float64 `json:"saving,omitempty"`
	// Discount 商品折扣 最小精度小数点后3位
	Discount float64 `json:"discount,omitempty"`
	// Comments 商品评论数
	Comments int `json:"comments,omitempty"`
	// Address 商品销售地址
	Address string `json:"address,omitempty"`
	// Feature 商品特色信息
	Feature string `json:"feature,omitempty"`
	// SalesPromotion 商品促销活动信息
	SalesPromotion string `json:"sales_promotion,omitempty"`
	// AgeV2 商品定向年龄 18-23, 24-30, 31-40, 41-49, 50+
	AgeV2 []string `json:"age_v2,omitempty"`
	// Gender 商品定向性别 0-不限, 1-男, 2-女
	Gender int `json:"gender,omitempty"`
	// ProvinceV2 商品定向省份
	ProvinceV2 []string `json:"province_v2,omitempty"`
	// CityV2 商品定向城市
	CityV2 []string `json:"city_v2,omitempty"`
	// Stock 商品库存状态 0-无, 1-有
	Stock int `json:"stock,omitempty"`
	// Status 商品投放状态 0-不可投放, 1-可投放
	Status int `json:"status,omitempty"`
	// CheckStatus 商品校验(接入)状态 0-成功, 1-失败
	CheckStatus int `json:"check_status,omitempty"`
	// CheckError 商品校验失败原因
	CheckError []string `json:"check_error,omitempty"`
	// OnlineTime 商品上线时间 毫秒时间戳
	OnlineTime int64 `json:"online_time,omitempty"`
	// OfflineTime 商品下线时间 毫秒时间戳
	OfflineTime int64 `json:"offline_time,omitempty"`
	// CreateTime 商品创建时间 毫秒时间戳
	CreateTime int64 `json:"create_time,omitempty"`
	// UpdateTime 商品更新时间 毫秒时间戳
	UpdateTime int64 `json:"update_time,omitempty"`
	// Slogan 商品广告语
	Slogan string `json:"slogan,omitempty"`
	// EcomProductInfo 电商行业商品特殊信息
	EcomProductInfo EcomProductInfoSneak `json:"ecom_product_info,omitempty"`
	// InsuranceProductInfo 保险行业商品特殊信息
	InsuranceProductInfo InsuranceProductInfoSneak `json:"insurance_product_info,omitempty"`
	// StackRoomProductInfo 小说行业商品特殊信息
	StackRoomProductInfo StackRoomProductInfoSneak `json:"stack_room_product_info,omitempty"`
	// AuVideoProductInfo 音视频行业商品特殊信息
	AuVideoProductInfo AuVideoProductInfoSneak `json:"au_video_product_info,omitempty"`
	// PharmaceuticalProductInfo 医药行业商品特殊信息
	PharmaceuticalProductInfo PharmaceuticalProductInfoSneak `json:"pharmaceutical_product_info,omitempty"`
	// RecruitmentProductInfo 招聘行业商品特殊信息
	RecruitmentProductInfo RecruitmentProductInfoSneak `json:"recruitment_product_info,omitempty"`
	// RegionalServiceProductInfo 区域服务行业商品特殊信息
	RegionalServiceProductInfo RegionalServiceProductInfoSneak `json:"regional_service_product_info,omitempty"`
	// CommunicationProductInfo 通信行业商品特殊信息
	CommunicationProductInfo CommunicationProductInfoSneak `json:"communication_product_info,omitempty"`
}

// VideoInfo 商品视频信息
type VideoInfo struct {
	// URL 视频URL
	URL string `json:"url,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// Duration 视频时长
	Duration int `json:"duration,omitempty"`
	// Ratio 视频码率
	Ratio int `json:"ratio,omitempty"`
}

// UrlParamSneak 商品链接信息
type UrlParamSneak struct {
	// PC PC端商品落地页URL
	PC string `json:"pc,omitempty"`
	// Mobile H5页面商品落地页URL
	Mobile string `json:"mobile,omitempty"`
	// AndroidApp Android应用商品直达吊起链接
	AndroidApp string `json:"android_app,omitempty"`
	// IOSApp IOS应用商品直达吊起链接
	IOSApp string `json:"ios_app,omitempty"`
	// UniversalLink IOS应用商品吊起ulink链接
	UniversalLink string `json:"universal_link,omitempty"`
}

// EcomProductInfoSneak 电商行业商品特殊信息
type EcomProductInfoSneak struct {
	// SpuID 标准商品单位ID
	SpuID string `json:"spu_id,omitempty"`
	// Bought 商品销量
	Bought int `json:"bought,omitempty"`
	// Comments 商品评论数
	Comments int `json:"comments,omitempty"`
	// Mark 商品评分
	Mark float64 `json:"mark,omitempty"`
	// Condition 商品新旧情况 全新, 翻新, 二手
	Condition string `json:"condition,omitempty"`
}

// InsuranceProductInfoSneak 保险行业商品特殊信息
type InsuranceProductInfoSneak struct {
	// Insurer 保险公司
	Insurer string `json:"insurer,omitempty"`
	// InsurerType 承保公司性质
	InsurerType string `json:"insurer_type,omitempty"`
	// Spokesperson 投放保险公司代言人
	Spokesperson string `json:"spokesperson,omitempty"`
	// InsuredMinAge 投保最小年龄
	InsuredMinAge int `json:"insured_min_age,omitempty"`
	// InsuredMaxAge 投保最大年龄
	InsuredMaxAge int `json:"insured_max_age,omitempty"`
	// PayoutAmount 保险金额
	PayoutAmount float64 `json:"payout_amount,omitempty"`
	// PayoutRatio 保险费率
	PayoutRatio float64 `json:"payout_ratio,omitempty"`
}

// StackRoomProductInfoSneak 小说行业商品特殊信息
type StackRoomProductInfoSneak struct {
	// AuthorID 作者ID
	AuthorID string `json:"author_id,omitempty"`
	// AuthorName 作者名
	AuthorName string `json:"author_name,omitempty"`
	// Words 字数
	Words int `json:"words,omitempty"`
	// Series 是否为连载书籍 0-否, 1-是
	Series int `json:"series,omitempty"`
	// Status 书籍状态 连载, 完本
	Status string `json:"status,omitempty"`
}

// AuVideoProductInfoSneak 音视频行业商品特殊信息
type AuVideoProductInfoSneak struct {
	// Category 视频类型
	Category string `json:"category,omitempty"`
	// Definition 视频清晰度
	Definition string `json:"definition,omitempty"`
	// Subtitles 视频字幕 0-无, 1-有
	Subtitles int `json:"subtitles,omitempty"`
}

// PharmaceuticalProductInfoSneak 医药行业商品特殊信息
type PharmaceuticalProductInfoSneak struct {
	// ApprovalNumber 药品批准文号
	ApprovalNumber string `json:"approval_number,omitempty"`
	// GenericName 通用名
	GenericName string `json:"generic_name,omitempty"`
	// Form 剂型
	Form string `json:"form,omitempty"`
	// DrugClass 药品类别
	DrugClass string `json:"drug_class,omitempty"`
}

// RecruitmentProductInfoSneak 招聘行业商品特殊信息
type RecruitmentProductInfoSneak struct {
	// Recruiter 招聘公司
	Recruiter string `json:"recruiter,omitempty"`
	// Position 招聘岗位
	Position string `json:"position,omitempty"`
	// Requirement 招聘要求
	Requirement string `json:"requirement,omitempty"`
	// Salary 薪资
	Salary string `json:"salary,omitempty"`
	// Benefits 福利
	Benefits []string `json:"benefits,omitempty"`
}

// RegionalServiceProductInfoSneak 区域服务行业商品特殊信息
type RegionalServiceProductInfoSneak struct {
	// ServiceProvider 服务提供商
	ServiceProvider string `json:"service_provider,omitempty"`
	// ServiceArea 服务区域
	ServiceArea []string `json:"service_area,omitempty"`
	// ServiceTime 服务时间
	ServiceTime string `json:"service_time,omitempty"`
	// Price 价格
	Price float64 `json:"price,omitempty"`
}

// CommunicationProductInfoSneak 通信行业商品特殊信息
type CommunicationProductInfoSneak struct {
	// Operator 运营商
	Operator string `json:"operator,omitempty"`
	// Plan 套餐
	Plan string `json:"plan,omitempty"`
	// MonthlyFee 月费用
	MonthlyFee float64 `json:"monthly_fee,omitempty"`
}
