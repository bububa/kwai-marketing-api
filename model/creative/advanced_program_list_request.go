package creative

import "encoding/json"

// AdvancedProgramListRequest 获取程序化创意2.0信息 API Request
type AdvancedProgramListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// UnitIDs 广告组ID;不超过一百个，当unit_ids参数不为空时，返回值total_count为0
	UnitIDs []int64 `json:"unit_ids,omitempty"`
	// PackageName 程序化创意包名称; 非空，0到100字符
	PackageName string `json:"package_name,omitempty"`
	// Status 程序化创意状态; -2：所有（包含已删除）、40：只包含已删除 不传：所有（不包含已删除）
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
}

// Url implement PostRequest interface
func (r AdvancedProgramListRequest) Url() string {
	return "v2/creative/advanced/program/list"
}

// Encode implement PostRequest interface
func (r AdvancedProgramListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
