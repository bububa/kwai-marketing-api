package creative

// BatchUpdateResponse 批量创建&修改创意 API Response
type BatchUpdateResponse struct {
	UpdateCreativeIDs []int64 `json:"update_creative_ids,omitempty"`
	AddCreativeIDs    []int64 `json:"add_creative_ids,omitempty"`
}
