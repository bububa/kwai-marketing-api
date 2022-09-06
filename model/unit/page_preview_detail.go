package unit

// PagePreviewDetail 程序化落地页组下的页面信息
type PagePreviewDetail struct {
	// PageID 落地页ID
	PageID string `json:"page_id,omitempty"`
	// ReviewStatus 落地页审核状态	1-审核中 2-审核通过 3-审核拒绝
	ReviewStatus int `json:"review_status,omitempty"`
	// ReviewDetail 落地页的审核拒绝理由
	ReviewDetail string `json:"review_detail,omitempty"`
	// URL 落地页URL信息
	URL string `json:"url,omitempty"`
}
