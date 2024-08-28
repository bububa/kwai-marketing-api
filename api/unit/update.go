package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/unit"
)

// Update 修改广告组
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateRequest) (uint64, error) {
	var resp unit.UpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
