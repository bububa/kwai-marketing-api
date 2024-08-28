package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// Create 创建广告组
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.CreateRequest) (uint64, error) {
	var resp unit.CreateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
