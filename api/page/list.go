package page

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/page"
)

// List 获取魔力建站落地页组信息列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *page.ListRequest) (*page.ListResponse, error) {
	var resp page.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
