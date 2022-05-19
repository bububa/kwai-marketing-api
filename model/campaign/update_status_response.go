package campaign

// UpdateStatusResponse 修改广告计划状态 API Response
type UpdateStatusResponse struct {
	// CampaignIDs 所有修改状态成功的计划id; 假如接口的入参 campaign_id传了值且修改状态成功，则此广告计划id也会包含在返回值campaign_ids里面。
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
}
