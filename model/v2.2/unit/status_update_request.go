package unit

import "encoding/json"

type StatusUpdateRequest struct {
	//advertiser_id	long	必填	广告主 ID	在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//unit_ids	long	必填	广告计划 ID
	UnitIds []uint64 `json:"unit_ids"`
	//1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status"`
}

// Url implement PostRequest interface
func (r StatusUpdateRequest) Url() string {
	return "v1/ad_unit/update/status"
}

// Encode implement PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
