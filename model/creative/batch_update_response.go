package creative

// BatchUpdateResponse 批量创建&修改创意 API Response
type BatchUpdateResponse struct {
	UpdateCreativeIDs []uint64 `json:"update_creative_ids,omitempty"`
	AddCreativeIDs    []uint64 `json:"add_creative_ids,omitempty"`
}
