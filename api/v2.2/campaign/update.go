package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/campaign"
)

// Update 更新广告计划
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (uint64, error) {
	var resp campaign.UpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CampaignID, nil
}
