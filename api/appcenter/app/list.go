package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// List 【应用中心】获取应用列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ListRequest) (*app.ListResponse, error) {
	var resp app.ListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
