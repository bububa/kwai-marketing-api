package comment

import "encoding/json"

// SetTopRequest 评论置顶 API Request
type SetTopRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentID 评论ID
	CommentID uint64 `json:"comment_id,omitempty"`
	// PhotoID 视频ID
	PhotoID uint64 `json:"photo_id,omitempty"`
}

// Url implement PostRequest interface
func (r SetTopRequest) Url() string {
	return "v1/comment/setTop"
}

// Encode implement PostRequest interface
func (r SetTopRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// SetTopResponse 评论置顶 API Response
type SetTopResponse struct {
	// CommentID 评论ID
	CommentID uint64 `json:"comment_id,omitempty"`
}
