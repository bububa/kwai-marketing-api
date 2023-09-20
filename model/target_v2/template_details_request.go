package target_v2

import (
	"encoding/json"
)

// TemplateDetailsRequest 查询定向模板接口 API Request
type TemplateDetailsRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// 模板ID列表（默认ID和分页参数二选一必传）
	TemplateID []uint64 `json:"template_id"`
	// Page 请求的页码，默认为 1（page和page_size要么都传，要么都不传；都不传时需要传模板ID）
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r TemplateDetailsRequest) Url() string {
	return "gw/dsp/target/template/details"
}

// Encode implement GetRequest interface
func (r TemplateDetailsRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
