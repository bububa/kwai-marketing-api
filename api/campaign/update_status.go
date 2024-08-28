package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/campaign"
)

// UpdateStatus 修改广告计划状态
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) ([]uint64, error) {
	var resp campaign.UpdateStatusResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.CampaignIDs, nil
}
