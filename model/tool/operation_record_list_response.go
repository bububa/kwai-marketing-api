package tool

// OperationRecordListResponse 账户操作记录信息查询
type OperationRecordListResponse struct {
	// TotalCount 总数
	TotalCount int64 `json:"total_count,omitempty"`
	// Details 操作记录
	Details []OperationRecord `json:"details,omitempty"`
}
