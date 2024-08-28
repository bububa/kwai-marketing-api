package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/campaign"
)

// Create 创建广告计划
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) {
	var resp campaign.CreateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CampaignID, nil
}
