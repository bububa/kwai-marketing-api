package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// AdvancedProgramCreate 创建程序化2.0创意
func AdvancedProgramCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramCreateRequest) (uint64, error) {
	var resp creative.AdvancedProgramCreateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
