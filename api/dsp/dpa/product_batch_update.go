package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// ProductBatchUpdate 更新商品
func ProductBatchUpdate(clt *core.SDKClient, accessToken string, req *dpa.ProductBatchUpdateRequest) ([]dpa.ProductUpdateResult, error) {
	var resp dpa.ProductBatchUpdateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.ProductEditResponses, nil
}
