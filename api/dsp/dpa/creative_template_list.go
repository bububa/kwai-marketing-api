package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// CreativeTemplateList 获取SDPA创意视频模板
func CreativeTemplateList(clt *core.SDKClient, accessToken string, req *dpa.CreativeTemplateListRequest) (*dpa.CreativeTemplateListResponse, error) {
	var resp dpa.CreativeTemplateListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
