package subpkg

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/subpkg"
)

// List 【应用中心】获取分包管理/回收站列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.ListRequest) (*subpkg.ListResponse, error) {
	var resp subpkg.ListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
