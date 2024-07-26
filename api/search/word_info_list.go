package search

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/search"
)

// WordInfoList 获取关键词列表
func WordInfoList(clt *core.SDKClient, accessToken string, req *search.WordInfoListRequest) (*search.WordInfoListResponse, error) {
	var resp search.WordInfoListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
