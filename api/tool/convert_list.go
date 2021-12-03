package tool

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/tool"
)

// ConvertList 获取可用的转化目标
func ConvertList(clt *core.SDKClient, accessToken string, req *tool.ConvertListRequest) (*tool.ConvertListResponse, error) {
	var resp tool.ConvertListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
