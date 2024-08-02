package dpa

// ProductInfo 商品信息
type ProductInfo struct {
	// IndustryID 商品所属行业ID
	// 1-电商, 3-保险, 5-书库, 8-音视频, 9-医药, 10-招聘, 11-区域服务, 12-通信
	IndustryID int `json:"industry_id,omitempty"`
	// SubIndustryID 商品类型ID(所属二级行业ID)
	// 书库: 1-网络小说, 2-漫画, 3-有声读物, 4-网赚; 音视频: 电影, 动漫, 纪录片, 综艺, 电视剧, 短剧, 音乐, 直播, 网赚, 其他
	SubIndustryID string `json:"sub_industry_id,omitempty"`
	// IndustryName 商品行业名称 1-电商, 3-保险, 5-书库, 8-音视频, 9-医药, 10-招聘, 11-区域服务, 12-通信
	IndustryName string `json:"industry_name,omitempty"`
	// SubIndustryName 商品类型名称(所属二级行业名称) 小说: 1-网络小说, 2-漫画, 3-有声读物, 4-网赚; 音视频: 电影, 动漫, 纪录片, 综艺, 电视剧, 短剧, 音乐, 直播, 网赚, 其他
	SubIndustryName string `json:"sub_industry_name,omitempty"`
	// LibraryID 商品库ID
	// 非XML商品库，商品库所属行业与商品所属行业需保持一致
	LibraryID uint64 `json:"library_id,omitempty"`
	// OuterID 商品第三方ID
	// 在商品库中保证唯一不可重复，长度小于48
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
	Video *VideoInfo `json:"video,omitempty"`
	// TargetURLs 商品链接信息
	TargetURLs *UrlParamSneak `json:"target_urls,omitempty"`
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
	// Value 商品原价
	// 最小精度小数点后3位
	Value float64 `json:"value,omitempty"`
	// Price 商品现价
	// 最小精度小数点后3位
	Price float64 `json:"price,omitempty"`
	// Saving 商品优惠价格
	// 最小精度小数点后3位
	Saving float64 `json:"saving,omitempty"`
	// Discount 商品折扣
	// 最小精度小数点后3位
	Discount float64 `json:"discount,omitempty"`
	// Comments 商品评论数
	Comments int `json:"comments,omitempty"`
	// Address 商品销售地址
	Address string `json:"address,omitempty"`
	// Feature 商品特色信息
	Feature string `json:"feature,omitempty"`
	// SalesPromotion 商品促销活动信息
	SalesPromotion string `json:"sales_promotion,omitempty"`
	// AgeV2 商品定向年龄
	// 18-23, 24-30, 31-40, 41-49, 50+
	AgeV2 []string `json:"age_v2,omitempty"`
	// Gender 商品定向性别
	// 0-不限, 1-男, 2-女
	Gender int `json:"gender,omitempty"`
	// ProvinceV2 商品定向省份
	ProvinceV2 []string `json:"province_v2,omitempty"`
	// CityV2 商品定向城市
	CityV2 []string `json:"city_v2,omitempty"`
	// Stock 商品库存状态
	// 0-无, 1-有
	Stock int `json:"stock,omitempty"`
	// Status 商品投放状态
	// 0-不可投放, 1-可投放
	Status int `json:"status,omitempty"`
	// CheckStatus
	// 商品校验(接入)状态 0-成功, 1-失败
	CheckStatus int `json:"check_status,omitempty"`
	// CheckError 商品校验失败原因
	CheckError []string `json:"check_error,omitempty"`
	// OnlineTime 商品上线时间
	// 毫秒时间戳
	OnlineTime int64 `json:"online_time,omitempty"`
	// OfflineTime 商品下线时间
	// 毫秒时间戳
	OfflineTime int64 `json:"offline_time,omitempty"`
	// CreateTime 商品创建时间
	// 毫秒时间戳
	CreateTime int64 `json:"create_time,omitempty"`
	// UpdateTime 商品更新时间
	// 毫秒时间戳
	UpdateTime int64 `json:"update_time,omitempty"`
	// Slogan 商品广告语
	Slogan string `json:"slogan,omitempty"`
	// EcomProductInfo 电商行业商品特殊信息
	EcomProductInfo *EcomProductInfoSneak `json:"ecom_product_info,omitempty"`
	// InsuranceProductInfo 保险行业商品特殊信息
	InsuranceProductInfo *InsuranceProductInfoSneak `json:"insurance_product_info,omitempty"`
	// StackRoomProductInfo 小说行业商品特殊信息
	StackRoomProductInfo *StackRoomProductInfoSneak `json:"stack_room_product_info,omitempty"`
	// AuVideoProductInfo 音视频行业商品特殊信息
	AuVideoProductInfo *AuVideoProductInfoSneak `json:"au_video_product_info,omitempty"`
	// PharmaceuticalProductInfo 医药行业商品特殊信息
	PharmaceuticalProductInfo *PharmaceuticalProductInfoSneak `json:"pharmaceutical_product_info,omitempty"`
	// RecruitmentProductInfo 招聘行业商品特殊信息
	RecruitmentProductInfo *RecruitmentProductInfoSneak `json:"recruitment_product_info,omitempty"`
	// RegionalServiceProductInfo 区域服务行业商品特殊信息
	RegionalServiceProductInfo *RegionalServiceProductInfoSneak `json:"regional_service_product_info,omitempty"`
	// CommunicationProductInfo 通信行业商品特殊信息
	CommunicationProductInfo *CommunicationProductInfoSneak `json:"communication_product_info,omitempty"`
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
	// Condition 商品新旧情况 全新, 翻新, 二手
	Condition string `json:"condition,omitempty"`
	// Mark 商品评分
	Mark float64 `json:"mark,omitempty"`
	// Bought 商品销量
	Bought int `json:"bought,omitempty"`
	// Comments 商品评论数
	Comments int `json:"comments,omitempty"`
}

// InsuranceProductInfoSneak 保险行业商品特殊信息
type InsuranceProductInfoSneak struct {
	// SpecialPayOut XXX医疗保险金
	SpecialPayOut *InsuranceSpecialPayOutInfoSneak `json:"special_pay_out,omitempty"`
	// Insurer 保险公司
	Insurer string `json:"insurer,omitempty"`
	// InsurerType 承保公司性质
	InsurerType string `json:"insurer_type,omitempty"`
	// HosAddress 医院范围
	HosAddress string `json:"hos_address,omitempty"`
	// StartTime 生效时间
	StartTime string `json:"start_time,omitempty"`
	// FreeAmount 免赔额
	FreeAmount string `json:"free_amount,omitempty"`
	// Spokesperson 投放保险公司代言人
	Spokesperson string `json:"spokesperson,omitempty"`
	// InAddress 保障地区
	InAddress string `json:"in_address,omitempty"`
	// HighestProtectionCost 最高保额
	HighestProtectionCost string `json:"highest_protection_cost,omitempty"`
	// PaymentTerm 缴费年限
	PaymentTerm string `json:"payment_term,omitempty"`
	// ComplimentaryInsurance 赠险
	// 有, 无
	ComplimentaryInsurance string `json:"complimentary_insurance,omitempty"`
	// PaymentMethod 缴费方式
	// 月缴, 年缴
	PaymentMethod string `json:"payment_method,omitempty"`
	// PaymentChannel 支付渠道
	// 支付宝, 微信, 支付宝或微信
	PaymentChannel string `json:"payment_channel,omitempty"`
	// Upgraded 链路中是否包含升级
	// 是, 否
	Upgraded string `json:"upgraded,omitempty"`
	// PayoutAmount 保险金额
	PayoutAmount float64 `json:"payout_amount,omitempty"`
	// PayoutRatio 保险费率
	PayoutRatio float64 `json:"payout_ratio,omitempty"`
	// TotalPrice 一次缴费价格
	TotalPrice float64 `json:"total_price,omitempty"`
	// FirstMonthPrice 首月价格
	FirstMonthPrice float64 `json:"first_month_price,omitempty"`
	// MonthPrice 次月价格
	MonthPrice float64 `json:"month_price,omitempty"`
	// NormalPayout 一般医疗保险金
	NormalPayout float64 `json:"normal_payout,omitempty"`
	// DiseasePayout 重疾保险金
	DiseasePayout float64 `json:"disease_payout,omitempty"`
	// InsuredMinAge 投保最小年龄
	InsuredMinAge int `json:"insured_min_age,omitempty"`
	// InsuredMaxAge 投保最大年龄
	InsuredMaxAge int `json:"insured_max_age,omitempty"`
	// Staging 是否分期
	// 1 - 分期 2- 不分期
	Staging int `json:"staging,omitempty"`
	// SocialIndemnity 是否有社保
	// 1-有社保, 2-无社保
	SocialIndemnity int `json:"social_indemnity,omitempty"`
	// AutoDeduction 是否自动扣费
	// 1-是, 2-否
	AutoDeduction int `json:"auto_deduction,omitempty"`
	// AutoRenew 是否自动续保
	// 1-是, 2-否
	AutoRenew int `json:"auto_renew,omitempty"`
	// Green 是否包含绿色就医服务
	// 1-是, 2-否
	Green int `json:"green,omitempty"`
	// InsurerTime 保障期限：单位(天)
	InsurerTime int `json:"insurer_time,omitempty"`
}

// InsuranceSpecialPayOutInfoSneak XXX医疗保险金
type InsuranceSpecialPayOutInfoSneak struct {
	// Name 保险金种类
	Name string `json:"name,omitempty"`
	// Payout 赔付金额（单位元）
	Payout float64 `json:"payout,omitempty"`
	// Wait 单位（天）
	Wait int `json:"wait,omitempty"`
}

// StackRoomProductInfoSneak 小说行业商品特殊信息
type StackRoomProductInfoSneak struct {
	// Author 作者
	Author string `json:"author,omitempty"`
	// ExtraTags 作品标签
	ExtraTags string `json:"extra_tags,omitempty"`
	// ChapterName1 章节一名称
	ChapterName1 string `json:"chapter_name1,omitempty"`
	// ChapterContent1 章节一内容
	ChapterContent1 string `json:"chapter_content1,omitempty"`
	// ChapterName2 章节二名称
	ChapterName2 string `json:"chapter_name2,omitempty"`
	// ChapterContent2 章节二内容
	ChapterContent2 string `json:"chapter_content2,omitempty"`
	// ChapterName3 章节三名称
	ChapterName3 string `json:"chapter_name3,omitempty"`
	// ChapterContent3 章节三内容
	ChapterContent3 string `json:"chapter_content3,omitempty"`
	// LongAndShortStories 书籍长短篇
	// 长篇书, 短篇书
	LongAndShortStories string `json:"long_and_short_stories,omitempty"`
	// BookCity 所属书城
	// 阅文,番茄, 点众, 阳光书城, 中文在线, 快看漫画, 网易, 掌中云, 悠书阁, 七猫, 番薯, 其他
	BookCity string `json:"book_city,omitempty"`
	// StartPaidChapter 起始付费章节
	StartPaidChapter string `json:"start_paid_chapter,omitempty"`
	// Theme 书籍分类
	// 详见https://docs.qingque.cn/s/home/eZQDJOM3GV-qQ4vNsEpaQ0qy2?identityId=F0DVUl94Gc
	Theme int `json:"theme,omitempty"`
	// Plot 题材
	// 详见https://docs.qingque.cn/s/home/eZQDJOM3GV-qQ4vNsEpaQ0qy2?identityId=F0DVUl94Gc
	Plot int `json:"plot,omitempty"`
	// Role 作品角色ID
	// 详见https://docs.qingque.cn/s/home/eZQDJOM3GV-qQ4vNsEpaQ0qy2?identityId=F0DVUl94Gc
	Role int `json:"role,omitempty"`
	// FictionStyle 作品风格ID
	// 详见https://docs.qingque.cn/s/home/eZQDJOM3GV-qQ4vNsEpaQ0qy2?identityId=F0DVUl94Gc
	FictionStyle int `json:"fiction_style,omitempty"`
	// Channel 作品频道ID
	// 详见https://docs.qingque.cn/s/home/eZQDJOM3GV-qQ4vNsEpaQ0qy2?identityId=F0DVUl94Gc
	Channel int `json:"channel,omitempty"`
	// OriginNation 作品地域ID
	// 详见https://docs.qingque.cn/s/home/eZQDJOM3GV-qQ4vNsEpaQ0qy2?identityId=F0DVUl94Gc
	OriginNation int `json:"origin_nation,omitempty"`
	// FirstStatus 作品完结状态
	// 1-未完结, 2-已完结
	FirstStatus int `json:"first_status,omitempty"`
	// PayStatus 作品付费类型
	// 1-付费, 2-免费
	PayStatus int `json:"pay_status,omitempty"`
	// IsAudioNovel 作品是否为有声小说
	// 1-有声, 2-非有声
	IsAudioNovel int `json:"is_audio_novel,omitempty"`
	// MinPayAmount 最低付费金额
	MinPayAmount int64 `json:"min_pay_amount,omitempty"`
	// FavoriteCount 书籍收藏数
	FavoriteCount int `json:"favorite_count,omitempty"`
}

// AuVideoProductInfoSneak 音视频行业商品特殊信息
type AuVideoProductInfoSneak struct {
	// Theme 作品题材
	// 详见https://docs.qingque.cn/s/home/eZQDk3ey4uwts_6HeBYVoxzTG?identityId=F0DVUl94Gc
	Theme string `json:"theme,omitempty"`
	// OriginNation 作品地域
	// 详见https://docs.qingque.cn/s/home/eZQDk3ey4uwts_6HeBYVoxzTG?identityId=F0DVUl94Gc
	OriginNation string `json:"origin_nation,omitempty"`
	// AlbumName 作品辑名称
	AlbumName string `json:"album_name,omitempty"`
	// PlatformAttributes 作品所处平台类型
	// 长视频, 短视频
	PlatformAttributes string `json:"platform_attributes,omitempty"`
	// Channel 作品男女频道	男频, 女频, 不限 (短剧类型音视频商品使用该字段)
	Channel string `json:"channel,omitempty"`
	// Plot 作品情节
	// 详见https://docs.qingque.cn/s/home/eZQDk3ey4uwts_6HeBYVoxzTG?identityId=F0DVUl94Gc (短剧类型音视频商品使用该字段)
	Plot string `json:"plot,omitempty"`
	// Copyright 版权方	自有版权请填写【主体名称】，分销请填写代投【主体名称】(短剧类型音视频商品使用该字段)
	Copyright string `json:"copyright,omitempty"`
	// Author 作品作者/导演
	Author []string `json:"author,omitempty"`
	// Starring 作品主角/主演
	Starring []string `json:"starring,omitempty"`
	// TalentID 作品达人ID
	// 短剧类型音视频商品使用该字段
	TalentID []string `json:"talent_id,omitempty"`
	// ReleaseDate 作品发布日期
	// 毫秒时间戳
	ReleaseDate int64 `json:"release_date,omitempty"`
	// AuthorFans 作者粉丝数
	AuthorFans int64 `json:"author_fans,omitempty"`
	// ImpressionNums 作品曝光量
	ImpressionNums int64 `json:"impression_nums,omitempty"`
	// PlayNums 作品播放量
	PlayNums int64 `json:"play_nums,omitempty"`
	// LikesNums 作品被喜欢/点赞次数
	LikesNums int64 `json:"likes_nums,omitempty"`
	// FavoriteUserNums 收藏作品用户数
	FavoriteUserNums int64 `json:"favorite_user_nums,omitempty"`
	// ForwardNums 作品转发数
	ForwardNums int64 `json:"forward_nums,omitempty"`
	// CommentNums 作品评论数
	CommentNums int64 `json:"comment_nums,omitempty"`
	// Score 作品评分
	Score float64 `json:"score,omitempty"`
	// AlbumSort 所处作品辑顺序
	AlbumSort int `json:"album_sort,omitempty"`
	// IsOnlineEarning 作品是否属于网赚
	// 1-是, 2-否 (非网赚类型音视频商品使用该字段)
	IsOnlineEarning int `json:"is_online_earning,omitempty"`
	// Episodes 剧类集数
	// 短剧类型音视频商品使用该字段
	Episodes int `json:"episodes,omitempty"`
	// EpisodeDuration 单集剧集分钟时长
	// 短剧类型音视频商品使用该字段
	EpisodeDuration int64 `json:"episode_duration,omitempty"`
	// PaidEpisodes 起始付费集数
	// 从第几集开始需要付费 (短剧类型音视频商品使用该字段)
	PaidEpisodes int `json:"paid_episodes,omitempty"`
	// UnitPrice 单集价格
	// 单位：元 (短剧类型音视频商品使用该字段)
	UnitPrice float64 `json:"unit_price,omitempty"`
	// AnnualVIPType 是否有年度会员
	// 1-有, 2-无 (短剧类型音视频商品使用该字段)
	AnnualVIPType int `json:"annual_vip_type,omitempty"`
	// SemiannualVIPType 是否有半年度会员
	// 1-有, 2-无 (短剧类型音视频商品使用该字段)
	SemiannualVIPType int `json:"semiannual_vip_type,omitempty"`
	// MaxCharge 最高充值档位
	// 单位：元 (短剧类型音视频商品使用该字段)
	MaxCharge float64 `json:"max_charge,omitempty"`
	// MinCharge 最低充值档位
	// 单位：元 (短剧类型音视频商品使用该字段)
	MinCharge float64 `json:"min_charge,omitempty"`
	// RecommendCharge 推荐充值档位
	// 单位：元 (短剧类型音视频商品使用该字段)
	RecommendCharge float64 `json:"recommend_charge,omitempty"`
}

// PharmaceuticalProductInfoSneak 医药行业商品特殊信息
type PharmaceuticalProductInfoSneak struct {
	// SpuID 标准商品单位ID
	SpuID string `json:"spu_id,omitempty"`
	// Tel 联系电话
	Tel string `json:"tel,omitempty"`
	// Nationality 所在国家
	Nationality string `json:"nationality,omitempty"`
	// Province 所在省份
	Province string `json:"province,omitempty"`
	// City 所在城市
	City string `json:"city,omitempty"`
	// District 所在区县
	District string `json:"district,omitempty"`
	// Region 所在商圈区域
	Region []string `json:"region,omitempty"`
	// ValueMin 商品原价区间最小值
	ValueMin float64 `json:"value_min,omitempty"`
	// ValueMax 商品原价区间最大值
	ValueMax float64 `json:"value_max,omitempty"`
	// PriceMin 商品现价区间最小值
	PriceMin float64 `json:"price_min,omitempty"`
	// PriceMax 商品现价区间最大值
	PriceMax float64 `json:"price_max,omitempty"`
	// SavingStartTime 减价开始时间
	// 毫秒时间戳
	SavingStartTime int64 `json:"saving_start_time,omitempty"`
	// SavingEndTime 减价结束时间
	// 毫秒时间戳
	SavingEndTime int64 `json:"saving_end_time,omitempty"`
	// ProductViewCount 商品详情页被打开次数
	ProductViewCount int64 `json:"product_view_count,omitempty"`
	// FavoriteCount 收藏商品用户数
	FavoriteCount int64 `json:"favorite_count,omitempty"`
	// FavourableCommentRate 商品好评率
	FavourableCommentRate int64 `json:"favourable_comment_rate,omitempty"`
	// Bought 商品销量
	Bought int `json:"bought,omitempty"`
	// Mark 商品评分
	Mark float64 `json:"mark,omitempty"`
}

// RecruitmentProductInfoSneak 招聘行业商品特殊信息
type RecruitmentProductInfoSneak struct {
	// JobType 工作类型
	// 全职, 兼职, 实习
	JobType string `json:"job_type,omitempty"`
	// PostType 岗位类型
	PostType string `json:"post_type,omitempty"`
	// SalaryOfYear 年薪
	SalaryOfYear string `json:"salary_of_year,omitempty"`
	// Salary 月薪
	Salary string `json:"salary,omitempty"`
	// Education 学历要求
	// 高中及以上, 技校及以上, 中专及以上, 大专及以上, 本科及以上, 硕士及以上, 博士及以上
	Education string `json:"education,omitempty"`
	// ServiceYear 工作年限要求
	// 1年以下, 1-2年, 3-5年, 6-7年, 8-10年, 10年以上
	ServiceYear string `json:"service_year,omitempty"`
	// Tel 联系电话
	Tel string `json:"tel,omitempty"`
	// Mail 联系邮箱
	Mail string `json:"mail,omitempty"`
	// Nationality 所在国家
	Nationality string `json:"nationality,omitempty"`
	// Province 所在省份
	Province string `json:"province,omitempty"`
	// City 所在城市
	City string `json:"city,omitempty"`
	// District 所在区县
	District string `json:"district,omitempty"`
	// Region 所在商圈区域
	Region []string `json:"region,omitempty"`
	// Benefits 福利待遇
	// 八险一金, 带薪休假, 包吃, 包住, 年底双薪, 周末双休, 交通补助, 加班补助, 饭补, 话补, 房补, 其他
	Benefits []string `json:"benefits,omitempty"`
	// HeadCount 招聘人数
	HeadCount int `json:"head_count,omitempty"`
}

// RegionalServiceProductInfoSneak 区域服务行业商品特殊信息
type RegionalServiceProductInfoSneak struct {
	// SpuID 外部商品附加ID
	SpuID string `json:"spu_id,omitempty"`
	// Tel 联系电话
	Tel string `json:"tel,omitempty"`
	// ProvinceID 所在省份ID
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	ProvinceID string `json:"province_id,omitempty"`
	// CityID 所在城市ID
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	CityID string `json:"city_id,omitempty"`
	// DistrictID 所在区县ID
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	DistrictID string `json:"district_id,omitempty"`
	// NationalityID 所在国家
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	NationalityID string `json:"nationality_id,omitempty"`
	// Province 所在省份
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	Province string `json:"province,omitempty"`
	// City 所在城市
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	City string `json:"city,omitempty"`
	// District 所在地区
	District string `json:"district,omitempty"`
	// Region 所在商圈区域
	Region []string `json:"region,omitempty"`
	// SavingStartTime 减价开始时间
	// 毫秒时间戳
	SavingStartTime int64 `json:"saving_start_time,omitempty"`
	// SavingEndTime 减价结束时间
	// 毫秒时间戳
	SavingEndTime int64 `json:"saving_end_time,omitempty"`
	// ProductViewCount 商品详情页被打开次数
	ProductViewCount int64 `json:"product_view_count,omitempty"`
	// FavoriteCount 收藏商品用户数
	FavoriteCount int64 `json:"favorite_count,omitempty"`
	// FavourableCommentRate 商品好评率
	FavourableCommentRate int64 `json:"favourable_comment_rate,omitempty"`
	// Bought 商品销量
	Bought int `json:"bought,omitempty"`
	// Mark 商品评分
	Mark float64 `json:"mark,omitempty"`
}

// CommunicationProductInfoSneak 通信行业商品特殊信息
type CommunicationProductInfoSneak struct {
	// TotalTraffic 总流量(单位GB)
	// 数字字符串
	TotalTraffic string `json:"total_traffic,omitempty"`
	// DirectionalTraffic 定向流量(单位GB)
	// 数字字符串
	DirectionalTraffic string `json:"directional_traffic,omitempty"`
	// GeneralTraffic 通用流量(单位GB)
	// 数字字符串
	GeneralTraffic string `json:"general_traffic,omitempty"`
	// ProvinceID 所在省份ID
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	ProvinceID string `json:"province_id,omitempty"`
	// CityID 所在城市ID
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	CityID string `json:"city_id,omitempty"`
	// DistrictID 所在区县ID
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	DistrictID string `json:"district_id,omitempty"`
	// NationalityID 所在国家
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	NationalityID string `json:"nationality_id,omitempty"`
	// Province 所在省份
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	Province string `json:"province,omitempty"`
	// City 所在城市
	// 详见接口https://ad.e.kuaishou.com/rest/openapi/v1/region/list
	City string `json:"city,omitempty"`
	// District 所在地区
	District string `json:"district,omitempty"`
	// Region 所在商圈区域
	Region []string `json:"region,omitempty"`
	// SavingStartTime 减价开始时间
	// 毫秒时间戳
	SavingStartTime int64 `json:"saving_start_time,omitempty"`
	// SavingStartTime 减价结束时间
	// 毫秒时间戳
	SavingEndTime int64 `json:"saving_end_time,omitempty"`
	// ProductViewCount 商品详情页被打开次数
	ProductViewCount int64 `json:"product_view_count,omitempty"`
	// FavoriteCount 收藏商品用户数
	FavoriteCount int64 `json:"favorite_count,omitempty"`
	// FavourableCommentRate 商品好评率
	FavourableCommentRate int64 `json:"favourable_comment_rate,omitempty"`
	// Bought 商品销量
	Bought int `json:"bought,omitempty"`
	// Mark 商品评分
	Mark float64 `json:"mark,omitempty"`
}
