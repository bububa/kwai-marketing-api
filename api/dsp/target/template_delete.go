package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateDelete 删除定向模板
func TemplateDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateDeleteRequest) (uint64, error) {
	var resp target.TemplateDeleteResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.TemplateID, nil
}
