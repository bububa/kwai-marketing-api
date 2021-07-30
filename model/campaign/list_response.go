package campaign

// ListResponse 获取广告计划信息 API Response
type ListResponse struct {
	// CampaignSubType 计划子类型; 4：DPA，5：SDPA
	CampaignSubType int `json:"campaign_sub_type,omitempty"`
}
