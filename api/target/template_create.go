package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/target"
)

// TemplateCreate 创建定向模板
func TemplateCreate(clt *core.SDKClient, accessToken string, req *target.TemplateCreateRequest) (*target.Template, error) {
	var resp target.Template
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
