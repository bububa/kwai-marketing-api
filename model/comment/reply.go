package comment

import "encoding/json"

// ReplyRequest 评论回复 API Request
type ReplyRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ReplyList 批量回复评论参数
	ReplyList []CommentReply `json:"reply_list,omitempty"`
}

// Url implement PostRequest interface
func (r ReplyRequest) Url() string {
	return "v1/comment/reply"
}

// Encode implement PostRequest interface
func (r ReplyRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ReplyResponse 评论回复 API Response
type ReplyResponse struct {
	// ReplyResultList
	ReplyResultList []ReplyResult `json:"reply_result_list,omitempty"`
}

type ReplyResult struct {
	// ReplyToCommentID 回复的评论 ID
	ReplyToCommentID uint64 `json:"reply_to_comment_id,omitempty"`
	// ReplyResult 回复结果
	// 1 成功；2 失败
	ReplyResult int `json:"reply_result,omitempty"`
}
