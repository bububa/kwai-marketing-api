package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateUpgradeList 查询待升级模板列表
func TemplateUpgradeList(clt *core.SDKClient, accessToken string, req *target.TemplateUpgradeListRequest) ([]target.TemplateUpgradeItem, error) {
	var resp []target.TemplateUpgradeItem
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
