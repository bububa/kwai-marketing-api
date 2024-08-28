package subpkg

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/subpkg"
)

// ReleaseList 获取新版分包发布列表【单元创编】
func ReleaseList(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.ReleaseListRequest) (*subpkg.ListResponse, error) {
	var resp subpkg.ListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
