package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/target"
)

// TemplateUpdate 修改定向模板
func TemplateUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateUpdateRequest) (*target.Template, error) {
	var resp target.Template
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
