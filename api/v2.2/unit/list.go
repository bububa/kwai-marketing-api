package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/unit"
)

// List 查询广告组
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.ListRequest) (*unit.ListResponse, error) {
	var resp unit.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
