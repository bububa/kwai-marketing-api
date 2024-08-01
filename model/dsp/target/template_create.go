package target

import (
	"github.com/bububa/kwai-marketing-api/model"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// TemplateCreateRequest 创建定向模板 API Request
type TemplateCreateRequest struct {
	// Target 定向信息
	Target *unit.Target `json:"target,omitempty"`
	// TemplateName 定向模板名称
	TemplateName string `json:"template_name,omitempty"`
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r TemplateCreateRequest) Url() string {
	return "gw/dsp/target/template/create"
}

// Encode implement GetRequest interface
func (r TemplateCreateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateCreateResponse 创建定向模板 API Response
type TemplateCreateResponse struct {
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}
