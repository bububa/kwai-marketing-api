package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// AdvancedProgramList 获取程序化创意2.0信息
func AdvancedProgramList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramListRequest) (*creative.AdvancedProgramListResponse, error) {
	var resp creative.AdvancedProgramListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
