package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/creative"
)

// List 查询广告创意
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ListRequest) (*creative.ListResponse, error) {
	var resp creative.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
