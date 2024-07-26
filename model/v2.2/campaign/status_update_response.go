package campaign

type StatusUpdateResponse struct {
	CampaignIds []uint64 `json:"campaign_ids"` // 所有修改状态成功的计划 id
}
