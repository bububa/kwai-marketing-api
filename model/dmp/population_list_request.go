package dmp

import (
	"encoding/json"
)

// PopulationListRequest 人群列表查询接口 API Request
type PopulationListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Status 人群包状态; 0, "计算中"；1, "已生效"；2, "已删除"；3, "推送中"；4, "已推送"；5, "计算失败"；6, "推送失败" ；7, "已失效"
	Status *int `json:"status,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页数，默认 20
	PageSize int `json:"page_size,omitempty"`
	//OrientationIds 人群包ID列表
	OrientationIds []uint64 `json:"orientation_ids,omitempty"`
}

// Url implement GetRequest interface
func (r PopulationListRequest) Url() string {
	return "v2/dmp/population/list"
}

// Encode implement GetRequest interface

func (r PopulationListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
