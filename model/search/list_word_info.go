package search

import "github.com/bububa/kwai-marketing-api/model"

// WordInfoListRequest 获取关键词列表
type WordInfoListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
}

// Url implement GetRequest interface
func (r WordInfoListRequest) Url() string {
	return "v2/word_info/list"
}

// Encode implement GetRequest interface
func (r WordInfoListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

type WordInfoListResponse struct {
	// TotalCount 数据总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 返回值详情
	Details []WordInfo `json:"details,omitempty"`
}

type WordInfo struct {
	WordInfoID   uint64 `json:"word_info_id,omitempty"`
	Word         string `json:"word,omitempty"`
	MatchType    int    `json:"match_type,omitempty"`
	ReviewStatus int    `json:"review_status,omitempty"`
	PutStatus    int    `json:"put_status,omitempty"`
	Status       int    `json:"status"`
}
