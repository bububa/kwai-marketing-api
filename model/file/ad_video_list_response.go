package file

// AdVideoListResponse 查询视频接口list接口 API Response
type AdVideoListResponse struct {
	// TotalCount 图片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 详情
	Details []Video `json:"details,omitempty"`
}
