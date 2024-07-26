package wordinfo

import "github.com/bububa/kwai-marketing-api/model"

// UpdateStatusRequest 修改关键词投放状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// WordInfoIDs ID不重复，最大数量20，关键词未删除
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
	// PutStatus 投放状态	1 - 投放，2 - 暂停
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateStatusRequest) Url() string {
	return "v2/word_info/update/status"
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateStatusResponse 修改关键词投放状态 API Response
type UpdateStatusResponse struct {
	// WordInfoIDs 修改成功关键词ID列表
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
}
