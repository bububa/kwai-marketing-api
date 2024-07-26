package comment

import "github.com/bububa/kwai-marketing-api/model"

// Comment 评论
type Comment struct {
	// CommentID 评论 ID
	CommentID uint64 `json:"comment_id,omitempty"`
	// RootCommentID 一级评论 ID
	RootCommentID uint64 `json:"root_comment_id,omitempty"`
	// CommentAuthorID 评论发布者 UserID
	CommentAuthorID uint64 `json:"comment_author_id,omitempty"`
	// PhotoID 视频 ID
	PhotoID model.Uint64 `json:"photo_id,omitempty"`
	// PhotoAuthorID 视频作者 UserID
	PhotoAuthorID uint64 `json:"photo_author_id,omitempty"`
	// CommentLevel 评论层级
	// 	1：一级评论，2：二级评论
	CommentLevel int `json:"comment_level,omitempty"`
	// CommentContent 评论内容
	CommentContent string `json:"comment_content,omitempty"`
	// FavNum 点赞数
	FavNum int64 `json:"fav_num,omitempty"`
	// PostTime 评论发布时间
	// 毫秒级时间戳
	PostTime int64 `json:"post_time,omitempty"`
	// NickName 评论发布者昵称
	NickName string `json:"nick_name,omitempty"`
	// ReplyStatus 回复状态
	// 1：未回复，2：已回复
	ReplyStatus int `json:"reply_status,omitempty"`
	// ShieldStatus 屏蔽状态
	// 1：未屏蔽，2：已屏蔽
	ShieldStatus int `json:"shield_status,omitempty"`
	// PhotoTags 视频标签
	PhotoTags []string `json:"photo_tags,omitempty"`
	// IsRootCommentForbid 一级评论是否被隐藏
	IsRootCommentForbid bool `json:"is_root_comment_forbid,omitempty"`
	// ShieldType 屏蔽类型
	// 1：评论内容关键词，3：手动屏蔽，4：社区自动屏蔽，5：用户昵称关键词，6：快手 ID
	ShieldType int `json:"shield_type,omitempty"`
	// EmotionURL 表情url
	EmotionURL string `json:"emotion_url,omitempty"`
	// IsTopComment 是否被置顶
	IsTopComment bool `json:"is_top_comment,omitempty"`
}

// CommentReply 评论回复
type CommentReply struct {
	// ReplyToCommentID 回复的评论 ID
	ReplyToCommentID uint64 `json:"reply_to_comment_id,omitempty"`
	// PhotoID 视频 ID
	PhotoID model.Uint64 `json:"photo_id,omitempty"`
	// PhotoAuthorID 视频作者 UserID
	PhotoAuthorID uint64 `json:"photo_author_id,omitempty"`
	// ReplyToUserID 被回复的用户 UserID
	ReplyToUserID uint64 `json:"reply_to_user_id,omitempty"`
	// ReplyContent 回复内容
	ReplyContent string `json:"reply_content,omitempty"`
}
