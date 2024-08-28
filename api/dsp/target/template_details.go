package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateDetails 查询定向模板
func TemplateDetails(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateDetailsRequest) (*target.TemplateDetailsResponse, error) {
	var resp target.TemplateDetailsResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
