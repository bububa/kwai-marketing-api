package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// Create 创建自定义创意
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreateRequest) (uint64, error) {
	var resp creative.CreateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CreativeID, nil
}
