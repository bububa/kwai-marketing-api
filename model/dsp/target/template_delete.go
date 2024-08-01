package target

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// TemplateDeleteRequest 删除定向模板 API Request
type TemplateDeleteRequest struct {
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateDeleteRequest) Url() string {
	return "gw/dsp/target/template/delete"
}

// Encode implement PostRequest interface
func (r TemplateDeleteRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateDeleteResponse API Response
type TemplateDeleteResponse struct {
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
}
