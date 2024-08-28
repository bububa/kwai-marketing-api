package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdImageList 查询图片信息list接口
func AdImageList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdImageListRequest) (*file.AdImageListResponse, error) {
	var resp file.AdImageListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
