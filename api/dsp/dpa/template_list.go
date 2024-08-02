package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// TemplateList 查询 DPA 模板信息
func TemplateList(clt *core.SDKClient, accessToken string, req *dpa.TemplateListRequest) (*dpa.TemplateListResponse, error) {
	var resp dpa.TemplateListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
