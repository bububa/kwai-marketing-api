package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/unit"
)

// List 获取广告组信息
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.ListRequest) (*unit.ListResponse, error) {
	var resp unit.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
