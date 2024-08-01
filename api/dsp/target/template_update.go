package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateUpdate 更新定向模板
func TemplateUpdate(clt *core.SDKClient, accessToken string, req *target.TemplateUpdateRequest) (uint64, error) {
	var resp target.TemplateUpdateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.TemplateID, nil
}
