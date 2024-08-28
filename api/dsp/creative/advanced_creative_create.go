package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// AdvancedCreativeCreate 创建程序化创意
func AdvancedCreativeCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeCreateRequest) (uint64, error) {
	var resp creative.AdvancedCreativeCreateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
