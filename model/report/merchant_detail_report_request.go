package report

import "encoding/json"

// MerchantDetailReportRequest 小店通转化数据APIRequest
type MerchantDetailReportRequest struct {
	// AdvertiserID 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// ViewType 数据类型：1：账户；2：广告计划；3：广告组；4：广告创意
	ViewType int `json:"view_type,omitempty"`
	// GroupType 汇总方式：1：天 （默认值）；2：小时 此时 start_date 只能是当天，end_date 只能是下一天
	GroupType int `json:"group_type,omitempty"`
	// StartDate 过滤筛选条件，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 过滤筛选条件，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// CampaignIDs 计划 id 列表，以英文逗号分割，查询计划转化数据 viewType=2 时必填，否则不用填写
	CampaignIDs []int64 `json:"campaign_ids,omitempty"`
	// UnitIDs 单元 id 列表，以英文逗号分割，查询单元转化数据 viewType=3 时必填，否则不用填写
	UnitIDs []int64 `json:"unit_ids,omitempty"`
	// CreativeIDs 创意 id 列表，以英文逗号分割，查询创意转化数据 viewType=4 时必填，否则不用填写
	CreativeIDs []int64 `json:"creative_ids,omitempty"`
	// ProgramedCreativeIDs 程序化创意 id 列表，以英文逗号分割，查询程序化创意转化数据 viewType=4 时必填，，否则不用填写
	ProgramedCreativeIDs []int64 `json:"programed_creative_ids,omitempty"`
	// Page 请求的页码，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认为 20，最大支持 2000
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest inteface
func (r MerchantDetailReportRequest) Url() string {
	return "v1/merchant/report/detail_report"
}

// Encode implement PostRequest interface
func (r MerchantDetailReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
