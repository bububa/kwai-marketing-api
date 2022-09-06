package unit

// PageGroupDetail 程序化落地页信息
type PageGroupDetail struct {
	// GroupID 程序化落地页组ID
	GroupID string `json:"group_id,omitempty"`
	// GroupName 程序化落地页组名称
	GroupName string `json:"group_name,omitempty"`
	// PagePreviewDetail 程序化落地页组下的页面信息
	PagePreviewDetail []PagePreviewDetail `json:"page_preview_detail,omitempty"`
}
