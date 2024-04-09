package unit

import "github.com/bububa/kwai-marketing-api/model"

// ListRequest 查询广告组 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	// 过滤筛选条件，若不传或传空则视为无限制条件
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 广告组ID
	// 过滤筛选条件，若不传或传空则视为无限制条件
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	// 过滤筛选条件，支持模糊搜索&精确查询
	UnitName string `json:"unit_name,omitempty"`
	// UnitIDs 广告组 ID 集
	// 不超过 100 个
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// Status 广告组状态
	// 过滤筛选条件；-2：不限(已删除)；10：广告组已删除；40：广告创意已删除.所有不包含已删除 其他值无效
	Status int `json:"status,omitempty"`
	// StartDate 开始时间
	// 与 end_date 同时传或同时不传；过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应 update_time 信息
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	// 与 start_date 同时传或同时不传；过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应 update_time 信息
	EndDate string `json:"end_date,omitempty"`
	// TimeFilterType 按创建时间或者更新时间进行筛选
	// 1.如传入此字段时不传"start_date"，与"end_date"字段，则不根据时间筛选。2.传入"start_date"，与"end_date"字段，且此字段为 1 时，按照创建时间进行筛选。3.传入"start_date"，与"end_date"字段，此字段不传，或传值为 0 时，则按照更新时间进行筛选
	TimeFilterType int `json:"time_filter_type,omitempty"`
	// Page 请求的页码数, 默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 请求的每页行数, 默认为 20
	PageSize int `json:"page_size,omitempty"`
	// ReviewStatusList 单元审核状态筛选
	// 1：待审核；2：审核通过；3：审核失败；7：待送审
	ReviewStatusList []int `json:"review_status_list,omitempty"`
	// PutStatusList 单元投放状态筛选
	// 1：投放；2：暂停；3：删除
	PutStatusList []int `json:"put_status_list,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "gw/dsp/unit/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ListResponse 获取广告组信息 API Response
type ListResponse struct {
	TotalCount int    `json:"total_count,omitempty"`
	Details    []Unit `json:"details,omitempty"`
}
