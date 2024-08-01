package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateUpgrade 模板升级
func TemplateUpgrade(clt *core.SDKClient, accessToken string, req *target.TemplateUpgradeRequest) (int64, error) {
	var resp target.TemplateUpgradeResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.TemplateCount, nil
}
