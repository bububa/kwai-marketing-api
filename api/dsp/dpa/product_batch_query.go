package dpa

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// ProductBatchQuery 获取商品列表
func ProductBatchQuery(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductBatchQueryRequest) (*dpa.ProductBatchQueryResponse, error) {
	var resp dpa.ProductBatchQueryResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
