package comment

import "encoding/json"

// CancelTopRequest 取消评论置顶 API Request
type CancelTopRequest struct {
	// AdvertiserID 告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentID 评论ID
	CommentID uint64 `json:"comment_id,omitempty"`
	// PhotoID 视频ID
	PhotoID uint64 `json:"photo_id,omitempty"`
}

// Url implement PostRequest interface
func (r CancelTopRequest) Url() string {
	return "v1/comment/cancelTop"
}

// Encode implement PostRequest interface
func (r CancelTopRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// CancelTopResponse 取消评论置顶 API Response
type CancelTopResponse struct {
	// CommentID 评论ID
	CommentID uint64 `json:"comment_id,omitempty"`
}
