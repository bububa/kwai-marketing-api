package file

// AdAppListResponse 获取应用列表 API Response
type AdAppListResponse struct {
	// TotalCount 返回条数
	TotalCount int `json:"total_count,omitempty"`
	// Details 列表
	Details []App `json:"details,omitempty"`
}
