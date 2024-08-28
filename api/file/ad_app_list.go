package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdAppList 获取应用列表
func AdAppList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdAppListRequest) (*file.AdAppListResponse, error) {
	var resp file.AdAppListResponse
	err := clt.GetOnBody(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
