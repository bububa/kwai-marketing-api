package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/campaign"
)

// List 查询广告计划
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.ListRequest) (*campaign.ListResponse, error) {
	var resp campaign.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
