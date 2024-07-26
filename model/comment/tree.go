package comment

import "encoding/json"

// TreeRequest 评论树查询 API Request
type TreeRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentID 评论 ID
	CommentID uint64 `json:"comment_id,omitempty"`
}

// Url implement PostRequest interface
func (r TreeRequest) Url() string {
	return "v1/comment/tree"
}

// Encode implement PostRequest interface
func (r TreeRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// TreeResponse 评论树查询 API Response
type TreeResponse struct {
	// RootCommentDetail 一级评论
	RootCommentDetail *Comment `json:"root_comment_detail,omitempty"`
	// ChildCommentDetailList 二级评论
	ChildCommentDetailList []Comment `json:"child_comment_detail_list,omitempty"`
}
