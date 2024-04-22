package wordinfo

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/wordinfo"
)

// List 获取关键词列表
func List(clt *core.SDKClient, accessToken string, req *wordinfo.ListRequest) (*wordinfo.ListResponse, error) {
	var resp wordinfo.ListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
