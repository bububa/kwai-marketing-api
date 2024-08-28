package tool

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// KeyFrame 获取可选的推荐封面
func KeyFrame(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.KeyFrameRequest) ([]string, error) {
	var resp []string
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
