package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/creative"
)

// AdvancedCreativeList 查询程序化广告创意
func AdvancedCreativeList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeListRequest) (*creative.ListAdvancedCreativeResponse, error) {
	var resp creative.ListAdvancedCreativeResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
