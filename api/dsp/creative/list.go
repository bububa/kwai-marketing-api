package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// List 查询自定义创意
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ListRequest) (*creative.ListResponse, error) {
	var resp creative.ListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
