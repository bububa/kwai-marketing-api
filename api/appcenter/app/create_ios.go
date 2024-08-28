package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// CreateIos 创建iOS应用
func CreateIos(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.CreateIosRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
