package creative

import "encoding/json"

// ListRequest 获取广告创意信息 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID; 过滤筛选条件，若不传或传空则视为无限制条件
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 广告组ID; 过滤筛选条件，若不传或传空则视为无限制条件
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativeID 广告创意ID; 过滤筛选条件，若不传或传空则视为无限制条件
	CreativeID string `json:"creative_id,omitempty"`
	// CreativeName 广告创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// CreativeIDs 广告创意ID集
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Status 广告创意状态; 过滤筛选条件；-1：不限，1：计划已暂停，3：计划超预算，6：余额不足，11：组审核中，12：组审核未通过，14：已结束，15：组已暂停，17：组超预算，19：未达投放时间，22：不在投放时段。41：审核中，42：审核未通过，46：已暂停，52：投放中，53：作品异常 -2：所有包含已删除 40：只包含已删除不传：所有不包含已删除 其他值无效
	Status int `json:"status,omitempty"`
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
	// ActionbarClickUrl 第三方点击按钮监测链接
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// ReviewStatusList 单元审核状态筛选	1：待审核；2：审核通过；3：审核失败；7：待送审
	ReviewStatusList []int `json:"review_status_list,omitempty"`
	// PutStatusList 单元投放状态筛选	1：投放；2：暂停；3：删除
	PutStatusList []int `json:"put_status_list,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "v1/creative/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
