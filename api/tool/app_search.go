package tool

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// AppSearch 获取可选的应用定向
func AppSearch(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.AppSearchRequest) (*tool.TargetingApp, error) {
	var resp tool.TargetingApp
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
