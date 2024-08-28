package dpa

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// CategoryList 获取商品库类目树
func CategoryList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.CategoryListRequest) (*dpa.CategoryListResponse, error) {
	var resp dpa.CategoryListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
