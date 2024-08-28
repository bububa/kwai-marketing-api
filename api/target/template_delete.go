package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/target"
)

// TemplateDelete 删除定向模板
func TemplateDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateDeleteRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
