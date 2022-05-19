package campaign

// ListResponse 获取广告计划信息 API Response
type ListResponse struct {
	TotalCount int       `json:"total_count"`
	Details    []Details `json:"details"`
}

type Details struct {
	Status            int     `json:"status"`
	CampaignId        uint64  `json:"campaign_id"`
	CampaignName      string  `json:"campaign_name"`
	PutStatus         int     `json:"put_status"`
	CreateChannel     int     `json:"create_channel"`
	DayBudget         int64   `json:"day_budget"`
	DayBudgetSchedule []int64 `json:"day_budget_schedule"`
	CampaignType      int     `json:"campaign_type"`
	CampaignSubType   int     `json:"campaign_sub_type"`
	CreateTime        string  `json:"create_time"`
	UpdateTime        string  `json:"update_time"`

	AdType int `json:"ad_type"`
}
