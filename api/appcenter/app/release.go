package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// Release 【应用中心】发布应用
func Release(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ReleaseRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
