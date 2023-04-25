package creative

import "encoding/json"

// 请求参数结构体
type AdvancedCreativeListRequest struct {
	AdvertiserID int64   `json:"advertiser_id"`          // 必填，广告主 ID，在获取 access_token 的时候返回
	UnitIDs      []int64 `json:"unit_ids,omitempty"`     // 选填，广告组 ID 集，不超过一百个
	PackageName  string  `json:"package_name,omitempty"` // 选填，程序化创意包名称，非空，0 到 100 字符
	Status       int     `json:"status,omitempty"`       // 选填，程序化创意状态，-2：所有（包含已删除）、40：只包含已删除，不传：所有（不包含已删除）
	StartDate    string  `json:"start_date,omitempty"`   // 选填，起始日期，格式如 yyyy-MM-dd
	EndDate      string  `json:"end_date,omitempty"`     // 选填，结束时期，格式如 yyyy-MM-dd
	// 1. 如传入此字段时不传"start_date"，与"end_date"字段，则不根据时间筛选。
	// 2. 传入"start_date"，与"end_date"字段，且此字段为 1 时，按照创建时间进行筛选。
	// 3. 传入"start_date"，与"end_date"字段，此字段不传，或传值为 0 时，则按照更新时间进行筛选
	TimeFilterType int   `json:"time_filter_type,omitempty"` // 选填，按创建时间或者更新时间进行筛选，
	Page           int   `json:"page,omitempty"`             // 选填，页数，默认为 1
	PageSize       int   `json:"page_size,omitempty"`        // 选填，每页行数，默认为 20
	PutStatusList  []int `json:"put_status_list,omitempty"`  // 选填，创意投放状态，1：投放；2：暂停；3：删除。备注：传了该参数会覆盖status参数筛选，因为二者是相同筛选项
}

// Url implement PostRequest interface
func (r AdvancedCreativeListRequest) Url() string {
	return "gw/dsp/advanced_creative/list"
}

// Encode implement PostRequest interface
func (r AdvancedCreativeListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
