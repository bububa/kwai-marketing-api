package creative

// CreativeTagAdviseResponse 创意标签填写建议 API Response
type CreativeTagAdviseResponse struct {
	// Industry 一级行业
	Industry string `json:"industry,omitempty"`
	// SecondIndustry 二级行业
	SecondIndustry string `json:"secondIndustry,omitempty"`
	// Tags 推荐标签
	Tags []string `json:"tags,omitempty"`
}
