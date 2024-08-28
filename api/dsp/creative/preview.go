package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// Preview 创意体验
func Preview(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.PreviewRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
