package unit

import "encoding/json"

// ListRequest 获取广告组信息 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID; 过滤筛选条件，若不传或传空则视为无限制条件
	CampaignID int64 `json:"campaign_id,omitempty"`
	// CampaignType 广告计划类型;过滤筛选条件；2 - 提升应用安装;3 - 获取电商下单;4 - 推广品牌活动;5 - 收集销售线索;
	CampaignType int `json:"campaign_type,omitempty"`
	// UnitID 广告组ID; 过滤筛选条件，若不传或传空则视为无限制条件
	UnitID int64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// UnitIDs 广告组ID集
	UnitIDs []int64 `json:"unit_ids,omitempty"`
	// Status 广告组状态;过滤筛选条件；-1：不限，1：计划已暂停，3：计划超预算，6：余额不足，11：审核中，12：审核未通过，14：已结束，15：已暂停，17：组超预算，19：未达投放时间，20：有效-2，22：不在投放时段。所有包含已删除10：只包含已删除不传：所有不包含已删除 其他值无效
	Status int `json:"status,omitempty"`
	// AppName 应用名称; 过滤筛选条件，支持模糊搜索
	AppName string `json:"app_name,omitempty"`
	// StartDate 开始时间;与end_date同时传或同时不传；过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应update_time信息
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间; 与start_date同时传或同时不传；过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应update_time信息
	EndDate string `json:"end_date,omitempty"`
	// TimeFilterType 按创建时间，还是更新时间进行筛选; 1.如传入此字段时不传"start_date"，与"end_date"字段，则不根据时间筛选。2.传入"start_date"，与"end_date"字段，且此字段为1时，按照创建时间进行筛选。3.传入"start_date"，与"end_date"字段，此字段不传，或传值为0时，则按照更新时间进行筛选
	TimeFilterType int `json:"time_filter_type,omitempty"`
	// Page 请求的页码数 默认为1
	Page int `json:"page,omitempty"`
	// PageSize 请求的每页行数; 默认为20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "v1/ad_unit/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
