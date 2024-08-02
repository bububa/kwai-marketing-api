package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// LibraryList 获取商品列表
func LibraryList(clt *core.SDKClient, accessToken string, req *dpa.LibraryListRequest) (*dpa.LibraryListResponse, error) {
	var resp dpa.LibraryListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
