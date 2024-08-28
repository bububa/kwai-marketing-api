package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// Detail 【应用中心】获取应用详情
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.DetailRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
