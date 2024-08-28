package live

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/live"
)

// ListLiveUser 获取主播列表
func ListLiveUser(ctx context.Context, clt *core.SDKClient, accessToken string, req *live.ListLiveUserRequest) (*live.ListLiveUserResponse, error) {
	var resp live.ListLiveUserResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
