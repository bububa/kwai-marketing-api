package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// UpdateStatus 修改创意状态
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) ([]uint64, error) {
	var resp creative.UpdateStatusResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.CreativeIDs, nil
}
