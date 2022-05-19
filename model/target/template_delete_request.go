package target

import "encoding/json"

// TemplateDeleteRequest 删除定向模板
type TemplateDeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TemplateID 定向模板id
	TemplateID string `json:"template_id,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateDeleteRequest) Url() string {
	return "v1/target/template/delete"
}

// Encode implement PostRequest interface
func (r TemplateDeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
