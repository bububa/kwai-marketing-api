package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// AdvancedCreativeList 查询程序化创意
func AdvancedCreativeList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeListRequest) (*creative.AdvancedCreativeListResponse, error) {
	var resp creative.AdvancedCreativeListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
