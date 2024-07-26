package wordinfo

import "github.com/bububa/kwai-marketing-api/model"

// CreateRequest 创建关键词 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 广告单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// WordInfos 关键词信息
	WordInfos []WordInfo `json:"word_infos,omitempty"`
}

// Url implement PostRequest interface
func (r CreateRequest) Url() string {
	return "v2/word_info/create"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CreateResponse 创建关键词 API Response
type CreateResponse struct {
	// ErrroList 添加失败关键词列表
	ErrroList []WordInfo `json:"error_list,omitempty"`
	// SuccessList 添加成功关键词列表
	SuccessList []WordInfo `json:"success_list,omitempty"`
}
