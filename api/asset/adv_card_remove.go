package asset

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/asset"
)

// AdvCardRemove 删除高级创意接口
func AdvCardRemove(ctx context.Context, clt *core.SDKClient, accessToken string, req *asset.AdvCardRemoveRequest) ([]uint64, error) {
	var resp asset.AdvCardRemoveResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.AdvCardID, nil
}
