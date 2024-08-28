package dpa

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// ProductBatchCreate 创建商品
func ProductBatchCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductBatchCreateRequest) ([]dpa.ProductUpdateResult, error) {
	var resp dpa.ProductBatchUpdateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.ProductEditResponses, nil
}
