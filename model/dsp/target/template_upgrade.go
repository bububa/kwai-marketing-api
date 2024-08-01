package target

import "github.com/bububa/kwai-marketing-api/model"

// TemplateUpgradeRequest 模板升级 API Request
type TemplateUpgradeRequest struct {
	// TemplateID 模板ID
	// 模板ID不为空时会指定升级该模板，当模板ID为空时，会更新账户下所有需要升级的模板
	TemplateID uint64 `json:"template_id,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateUpgradeRequest) Url() string {
	return "gw/dsp/target/template/upgrade"
}

// Encode implement PostRequest interface
func (r TemplateUpgradeRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateUpgradeResponse 模板升级 API Response
type TemplateUpgradeResponse struct {
	// TemplateCount 升级了的模板总数
	TemplateCount int64 `json:"template_count,omitempty"`
}
