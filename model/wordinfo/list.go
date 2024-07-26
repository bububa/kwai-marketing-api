package wordinfo

import "github.com/bububa/kwai-marketing-api/model"

// ListRequest 获取关键词列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 推广单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "v2/word_info/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ListResponse 获取关键词列表 API Response
type ListResponse struct {
	TotalCount int        `json:"total_count,omitempty"`
	Details    []WordInfo `json:"details,omitempty"`
}
