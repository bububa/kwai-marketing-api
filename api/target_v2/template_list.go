package target_v2

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/target_v2"
)

// TemplateList 查询定向模板接口
func TemplateDetails(clt *core.SDKClient, accessToken string, req *target_v2.TemplateDetailsRequest) (*target_v2.TemplateDetailsResponse, error) {
	var resp target_v2.TemplateDetailsResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
