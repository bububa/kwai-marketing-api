package campaign

import "encoding/json"

type ListRequest struct {
	AdvertiserID   uint64   `json:"advertiser_id"`              // 广告主 ID，在获取 access_token 的时候返回，必填
	CampaignID     uint64   `json:"campaign_id,omitempty"`      // 广告计划 ID，过滤筛选条件，若不传或传空则视为无限制条件
	CampaignName   string   `json:"campaign_name,omitempty"`    // 计划名称，过滤筛选条件，若不传或传空则视为无限制条件
	CampaignIDs    []uint64 `json:"campaign_ids,omitempty"`     // 广告计划 ID 集，不超过 200 个
	PutStatusList  int      `json:"put_status_list,omitempty"`  // 计划投放状态筛选，1：投放；2：暂停；3：删除。传了该参数会覆盖 status 参数筛选，因为二者是相同筛选项
	Status         int      `json:"status,omitempty"`           // 计划状态，过滤筛选条件；1：广告计划已暂停；4：有效；5：广告计划已删除； -2 不限
	StartDate      string   `json:"start_date,omitempty"`       // 开始时间，格式为"yyyy-MM-dd"，与 end_date 同时传或同时不传；过滤筛选条件，参数值对应 update_time 信息
	EndDate        string   `json:"end_date,omitempty"`         // 结束时间，格式为"yyyy-MM-dd"，与 start_date 同时传或同时不传；过滤筛选条件，参数值对应 update_time 信息
	TimeFilterType int      `json:"time_filter_type,omitempty"` // 按创建时间，还是更新时间进行筛选。1.如传入此字段时不传"start_date"，与"end_date"字段，则不根据时间筛选。2.传入"start_date"，与"end_date"字段，且此字段为 1 时，按照创建时间进行筛选。3.传入"start_date"，与"end_date"字段，此字段不传，或传值为 0 时，则按照更新时间进行筛选
	Page           int      `json:"page,omitempty"`             // 请求的页码数，默认为 1
	PageSize       int      `json:"page_size,omitempty"`        // 请求的每页行数，默认为 20
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "gw/dsp/campaign/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
