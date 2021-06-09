package advertiser

type Info struct {
	UserID              int64  `json:"user_id,omitempty"`               // 账户快手ID
	CorporationName     string `json:"corporation_name,omitempty"`      // 公司名称
	UserName            string `json:"user_name,omitempty"`             // 快手昵称
	IndustryID          int64  `json:"industry_id,omitempty"`           // 二级行业 id
	IndustryName        string `json:"industry_name,omitempty"`         // 二级行业名称
	PrimaryIndustryID   int64  `json:"primary_industry_id,omitempty"`   // 一级行业 id
	PrimaryIndustryName string `json:"primary_industry_name,omitempty"` // 一级行业名称
}
