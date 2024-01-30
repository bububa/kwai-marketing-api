package creative

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// UpdateStatusRequest 修改创意状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// CreativeID 广告创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeIDs 与原有的creative_id字段可以同时填，也可以只填一个; 1.传入的创意id不得重复，且至少有一个;传入的创意id总数最多20个。2.put_status为3时，会删除所有创意
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// PutStatus 操作码; 1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateStatusRequest) Url() string {
	return "v1/creative/update/status"
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// UpdateStatusResponse 修改创意状态 API Response
type UpdateStatusResponse struct {
	// CreativeIDs 所有修改状态成功的创意id; 假如接口的入参 creative_id传了值且修改状态成功，则此创意id也会包含在返回值creative_ids里面
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
}
