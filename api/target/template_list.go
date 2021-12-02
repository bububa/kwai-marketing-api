package target

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/target"
)

// TemplateList 查询定向模板接口
func TemplateList(clt *core.SDKClient, accessToken string, req *target.TemplateListRequest) (*target.TemplateListResponse, error) {
	var resp target.TemplateListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
