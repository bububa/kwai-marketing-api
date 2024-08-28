package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// ReleaseList 获取新版应用发布列表【单元创编】
func ReleaseList(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ReleaseListRequest) (*app.ListResponse, error) {
	var resp app.ListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
