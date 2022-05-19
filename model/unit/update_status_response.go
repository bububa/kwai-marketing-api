package unit

// UpdateStatusResponse 修改广告组状态 API Response
type UpdateStatusResponse struct {
	// UnitIDs 所有修改状态成功的广告组id; 假如接口的入参 unit_id传了值且修改状态成功，则此广告组id也会包含在返回值unit_ids里面。
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
}
