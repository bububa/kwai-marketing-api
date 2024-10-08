package live

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/live"
)

// ListJingleBell 获取直播间小铃铛列表
func ListJingleBell(ctx context.Context, clt *core.SDKClient, accessToken string, req *live.ListJingleBellRequest) (*live.ListJingleBellResponse, error) {
	var resp live.ListJingleBellResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
