package target

import "encoding/json"

// TemplateUpdateRequest 修改定向模板 API Request
type TemplateUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TemplateID 定向模板id
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateName 定向模板名称; 名字需要唯一
	TemplateName string `json:"template_name,omitempty"`
	// Target 定向信息
	Target *Target `json:"target,omitempty"`
	// BehaviorInterest 行为兴趣定向
	BehaviorInterest *BehaviorInterest `json:"behavior_interest,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateUpdateRequest) Url() string {
	return "v1/target/template/update"
}

// Encode implement PostRequest interface
func (r TemplateUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
