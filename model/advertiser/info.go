package advertiser

// Info 广告主信息
type Info struct {
	// UserID 账户快手ID
	UserID uint64 `json:"user_id,omitempty"`
	// CorporationName 公司名称
	CorporationName string `json:"corporation_name,omitempty"`
	// UserName 快手昵称
	UserName string `json:"user_name,omitempty"`
	// IndustryID 二级行业 id
	IndustryID uint64 `json:"industry_id,omitempty"`
	// IndustryName 二级行业名称
	IndustryName string `json:"industry_name,omitempty"`
	// PrimaryIndustryID 一级行业 id
	PrimaryIndustryID uint64 `json:"primary_industry_id,omitempty"`
	// PrimaryIndustryName 一级行业名称
	PrimaryIndustryName string `json:"primary_industry_name,omitempty"`
}
