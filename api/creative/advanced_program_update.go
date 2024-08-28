package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// AdvancedProgramUpdate 修改程序化2.0创意
func AdvancedProgramUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramUpdateRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
