package lp

// ConsultListResponse 获取可用咨询组件列表 API Response
type ConsultListResponse struct {
	// TotalCount 总共条数
	TotalCount int `json:"total_count,omitempty"`
	// Details json array
	Details []Consult `json:"details,omitempty"`
}
