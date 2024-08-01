package target

import (
	"github.com/bububa/kwai-marketing-api/model"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// TemplateUpdateRequest 更新定向模板 API Request
type TemplateUpdateRequest struct {
	// Target 定向信息
	Target *unit.Target `json:"target,omitempty"`
	// TemplateName 定向模板名称
	TemplateName string `json:"template_name,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r TemplateUpdateRequest) Url() string {
	return "gw/dsp/target/template/update"
}

// Encode implement GetRequest interface
func (r TemplateUpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateUpdateResponse 更新定向模板 API Response
type TemplateUpdateResponse struct {
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}
