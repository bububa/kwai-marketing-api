package target

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateDetails 查询定向模板
func TemplateDetails(clt *core.SDKClient, accessToken string, req *target.TemplateDetailsRequest) (*target.TemplateDetailsResponse, error) {
	var resp target.TemplateDetailsResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
