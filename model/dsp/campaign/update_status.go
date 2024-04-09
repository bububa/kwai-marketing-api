package campaign

import "github.com/bububa/kwai-marketing-api/model"

// UpdateStatusRequest 修改广告计划状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划 ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignIDs 与原有的 campaign_id 字段可以同时填，也可以只填一个
	// 1.传入的计划 id 不得重复，且至少有一个;传入的 campaign_id 总数最多 20 个。2.put_status 为 3 时，会删除广告计划，和计划下的所有广告组，程序化创意，创意
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// PutStatus 操作码
	// 1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateStatusRequest) Url() string {
	return "v1/campaign/update/status"
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateStatusResponse 修改广告计划状态 API Response
type UpdateStatusResponse struct {
	// CampaignIDs 所有修改状态成功的计划 id
	// 假如接口的入参 campaign_id 传了值且修改状态成功，则此广告计划 id 也会包含在返回值 campaign_ids 里面。
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
}
