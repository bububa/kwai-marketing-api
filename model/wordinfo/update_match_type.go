package wordinfo

import "github.com/bububa/kwai-marketing-api/model"

// UpdateMatchTypeRequest 修改关键词匹配方式 API Request
type UpdateMatchTypeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// WordInfoIDs ID不重复，最大数量20，关键词未删除
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
	// MatchType 1 - 精确匹配，2 - 短语匹配，3 - 广泛匹配
	MatchType int `json:"match_type,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateMatchTypeRequest) Url() string {
	return "v2/word_info/update/match_type"
}

// Encode implement PostRequest interface
func (r UpdateMatchTypeRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateMatchTypeResponse 修改关键词匹配方式 API Response
type UpdateMatchTypeResponse struct {
	// WordInfoIDs 修改成功关键词ID列表
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
}
