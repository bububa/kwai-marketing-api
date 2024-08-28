package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// AdvancedCreativeUpdate 修改程序化创意
func AdvancedCreativeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeUpdateRequest) (uint64, error) {
	var resp creative.AdvancedCreativeUpdateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
