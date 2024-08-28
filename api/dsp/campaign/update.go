package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/campaign"
)

// Update 修改广告计划
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (uint64, error) {
	var resp campaign.UpdateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CampaignID, nil
}
