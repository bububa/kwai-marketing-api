package campaign

import "github.com/bububa/kwai-marketing-api/model"

// StatusUpdateRequest 广告计划状态更新 API Request
type StatusUpdateRequest struct {
	// advertiser_id	long	必填	广告主 ID	在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// unit_ids	long	必填	广告计划 ID
	CampaignIds []uint64 `json:"campaign_ids,omitempty"`
	// 1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r StatusUpdateRequest) Url() string {
	return "v1/campaign/update/status"
}

// Encode implement PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// StatusUpdateResponse 广告计划状态更新 API Response
type StatusUpdateResponse struct {
	// CampaignIds 所有修改状态成功的计划 id
	CampaignIds []uint64 `json:"campaign_ids,omitempty"`
}
