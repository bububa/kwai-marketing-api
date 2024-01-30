package creative

import "github.com/bububa/kwai-marketing-api/model"

// ListRequest 查询自定义创意 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划 ID
	// 过滤筛选条件，若不传或传空则视为无限制条件
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 广告组ID
	// 过滤筛选条件，若不传或传空则视为无限制条件
	UnitID uint64 `json:"unit_id,omitempty"`
	// CreativeID 创意 ID
	// 过滤筛选条件，若不传或传空则视为无限制条件
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeName 创意名称
	// 过滤筛选条件，支持模糊搜索 精确查询
	CreativeName string `json:"creative_name,omitempty"`
	// CreativeIDs 广告创意 ID 集
	// 不超过 100 个
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Status 广告创意状态
	// 过滤筛选条件； -2:不限,40:只包含已删除 不传：所有不包含已删除 其他值无效
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
	// PutStatusList 创意投放状态
	// 1：投放；2：暂停；3：删除。备注：传了该参数会覆盖status参数筛选，因为二者是相同筛选项
	PutStatusList []int `json:"put_status_list,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "gw/dsp/creative/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ListResponse 查询自定义创意 API Response
type ListResponse struct {
	TotalCount int        `json:"total_count,omitempty"`
	Details    []Creative `json:"details,omitempty"`
}
