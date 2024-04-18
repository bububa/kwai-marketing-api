package comment

import "encoding/json"

// ListRequest 评论列表数据查询 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CommentContent 	评论内容
	CommentContent string `json:"comment_content,omitempty"`
	// ReplyStatus 回复状态
	// 1：未回复，2：已回复
	ReplyStatus int `json:"reply_status,omitempty"`
	// CommentLevel 评论层级
	// 1：一级评论，2：二级评论
	CommentLevel int `json:"comment_level,omitempty"`
	// PostStartTime 评论发布起始时间
	// 与 post_time_end 同时传或同时不传；过滤筛选条件，毫秒级时间戳
	PostStartTime int64 `json:"post_time_start,omitempty"`
	// PostEndTime 评论发布结束时间
	// 与 post_time_start 同时传或同时不传；过滤筛选条件，毫秒级时间戳
	PostEndTime int64 `json:"post_time_end,omitempty"`
	// ShieldStatus 隐藏状态
	// 1：未隐藏，2：已隐藏
	ShieldStatus int `json:"shield_status,omitempty"`
	// PhotoQueryValue 视频搜索参数
	// 可输入视频 ID 或视频名称进行查询
	PhotoQueryValue string `json:"photo_query_value,omitempty"`
	// PhotoTag 视频标签
	PhotoTag string `json:"photo_tag,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页条数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "v1/comment/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ListResponse 评论列表数据查询 API Response
type ListResponse struct {
	// Total 	符合条件的记录总数
	Total int `json:"total,omitempty"`
	// Details 	评论列表
	Details []Comment `json:"details,omitempty"`
}
