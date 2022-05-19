package unit

import "encoding/json"

// UpdateStatusRequest 修改广告组状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitIDs 与原有的unit_id字段可以同时填，也可以只填一个; 1.传入的广告组id不得重复，且至少有一个;传入的广告组id总数最多20个。2.put_status为3时，会删除广告组，广告组下的所有程序化创意，创意
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// PutStatus 操作码; 1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateStatusRequest) Url() string {
	return "v1/ad_unit/update/status"
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
