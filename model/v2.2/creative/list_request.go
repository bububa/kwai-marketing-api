package creative

import "encoding/json"

// 请求参数结构体
type ListRequest struct {
	AdvertiserID   uint64   `json:"advertiser_id"`              // 广告主 ID，在获取 access_token 的时候返回
	CampaignID     uint64   `json:"campaign_id,omitempty"`      // 广告计划 ID，过滤筛选条件，若不传或传空则视为无限制条件
	UnitID         uint64   `json:"unit_id,omitempty"`          // 广告组 ID，过滤筛选条件，若不传或传空则视为无限制条件
	CreativeID     uint64   `json:"creative_id,omitempty"`      // 广告创意 ID，过滤筛选条件，若不传或传空则视为无限制条件
	CreativeName   string   `json:"creative_name,omitempty"`    // 广告创意名称，过滤筛选条件，支持模糊搜索和精确查询
	CreativeIDs    []uint64 `json:"creative_ids,omitempty"`     // 广告创意 ID 集，不超过 100 个
	Status         int      `json:"status,omitempty"`           // 广告创意状态，过滤筛选条件，-2:不限,40:只包含已删除，不传：所有不包含已删除，其他值无效
	StartDate      string   `json:"start_date,omitempty"`       // 开始时间，过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应 update_time 信息
	EndDate        string   `json:"end_date,omitempty"`         // 结束时间，过滤筛选条件，格式为"yyyy-MM-dd"，参数值对应 update_time 信息
	TimeFilterType int      `json:"time_filter_type,omitempty"` // 按创建时间或者更新时间进行筛选，1.如传入此字段时不传"start_date"，与"end_date"字段，则不根据时间筛选。2.传入"start_date"，与"end_date"字段，且此字段为 1 时，按照创建时间进行筛选。3.传入"start_date"，与"end_date"字段，此字段不传，或传值为 0 时，则按照更新时间进行筛选
	Page           int      `json:"page,omitempty"`             // 请求的页码数，默认为 1
	PageSize       int      `json:"page_size,omitempty"`        // 请求的每页行数，默认为 20
	PutStatusList  []int    `json:"put_status_list,omitempty"`  // 创意投放状态，1：投放；2：暂停；3：删除。备注：传了该参数会覆盖 status 参数筛选，因为二者是相同筛选项
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "gw/dsp/creative/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
