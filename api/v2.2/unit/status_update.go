package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/unit"
)

// Create 创建广告组
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.StatusUpdateRequest) ([]uint64, error) {
	var resp unit.StatusUpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.UnitIds, nil
}
