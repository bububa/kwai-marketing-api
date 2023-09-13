package unit

type StatusUpdateResponse struct {
	UnitIds []uint64 `json:"unit_ids"` // 所有修改状态成功的广告组 id
}
