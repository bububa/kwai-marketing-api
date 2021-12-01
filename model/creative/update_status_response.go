package creative

// UpdateStatusResponse 修改创意状态
type UpdateStatusResponse struct {
	// CreativeIDs 所有修改状态成功的创意id; 假如接口的入参 creative_id传了值且修改状态成功，则此创意id也会包含在返回值creative_ids里面
	CreativeIDs []int64 `json:"creative_ids,omitempty"`
}
