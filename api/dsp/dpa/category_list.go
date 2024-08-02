package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// CategoryList 获取商品库类目树
func CategoryList(clt *core.SDKClient, accessToken string, req *dpa.CategoryListRequest) (*dpa.CategoryListResponse, error) {
	var resp dpa.CategoryListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
