package tool

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// ConvertList 获取可用的转化目标
func ConvertList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.ConvertListRequest) (*tool.ConvertListResponse, error) {
	var resp tool.ConvertListResponse
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
