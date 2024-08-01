package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateUnitSync 定向模板同步
func TemplateUnitSync(clt *core.SDKClient, accessToken string, req *target.TemplateUnitSyncRequest) (*target.TemplateUnitSyncResponse, error) {
	var resp target.TemplateUnitSyncResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
