package creative

import "github.com/bububa/kwai-marketing-api/model"

// AdvancedCreativeListRequest 查询程序化创意 API Request
type AdvancedCreativeListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitIDs 广告组 ID 集
	// 不超过一百个
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// PackageName 程序化创意包名称
	PackageName string `json:"package_name,omitempty"`
	// Status 程序化创意状态
	// -2：所有（包含已删除）、40：只包含已删除 不传：所有（不包含已删除）
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
	// PutStatusAdvancedCreativeList 创意投放状态
	// 1：投放；2：暂停；3：删除。备注：传了该参数会覆盖status参数筛选，因为二者是相同筛选项
	PutStatusAdvancedCreativeList []int `json:"put_status_list,omitempty"`
}

// Url implement PostRequest interface
func (r AdvancedCreativeListRequest) Url() string {
	return "gw/dsp/advanced_creative/list"
}

// Encode implement PostRequest interface
func (r AdvancedCreativeListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdvancedCreativeListResponse 查询程序化创意 API Response
type AdvancedCreativeListResponse struct {
	TotalCount int                `json:"total_count,omitempty"`
	Details    []AdvancedCreative `json:"details,omitempty"`
}
