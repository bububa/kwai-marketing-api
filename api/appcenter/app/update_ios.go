package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// UpdateIos 更新iOS应用
func UpdateIos(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.UpdateIosRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
