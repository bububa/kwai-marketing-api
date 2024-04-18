package comment

import "encoding/json"

// ShieldRequest 评论屏蔽 API Request
type ShieldRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ShieldList 批量屏蔽评论参数，单次请求最多屏蔽50条评论
	ShieldList []ShieldComment `json:"shield_list,omitempty"`
}

// SheildComment 屏蔽评论参数
type ShieldComment struct {
	// CommentID 评论ID
	CommentID uint64 `json:"comment_id,omitempty"`
	// PhotoID 视频 ID
	PhotoID uint64 `json:"photo_id,omitempty"`
}

// Url implement PostRequest interface
func (r ShieldRequest) Url() string {
	return "v1/comment/shield"
}

// Encode implement PostRequest interface
func (r ShieldRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
