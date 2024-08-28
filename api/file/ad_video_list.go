package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoList 查询视频信息list接口
func AdVideoList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoListRequest) (*file.AdVideoListResponse, error) {
	var resp file.AdVideoListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
