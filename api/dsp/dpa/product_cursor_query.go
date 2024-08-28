package dpa

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// ProductCursorQuery 获取商品列表(游标)
func ProductCursorQuery(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductCursorQueryRequest) (*dpa.ProductCursorQueryResponse, error) {
	var resp dpa.ProductCursorQueryResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
