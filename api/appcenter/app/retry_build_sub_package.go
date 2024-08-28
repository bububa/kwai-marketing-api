package app

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// RetryBuildSubPackage 【应用中心】分包失败，重新构建
func RetryBuildSubPackage(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.RetryBuildSubPackageRequest) (int, error) {
	var resp app.RetryBuildSubPackageResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.RetryCnt, nil
}
