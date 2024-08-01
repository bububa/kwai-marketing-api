package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateCreate 创建定向模板
func TemplateCreate(clt *core.SDKClient, accessToken string, req *target.TemplateCreateRequest) (uint64, error) {
	var resp target.TemplateCreateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.TemplateID, nil
}
