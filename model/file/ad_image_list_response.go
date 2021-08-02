package file

// AdImageListResponse 查询图片接口list接口 API Response
type AdImageListResponse struct {
	// TotalCount 图片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 详情
	Details []Image `json:"details,omitempty"`
}
