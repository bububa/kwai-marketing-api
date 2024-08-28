package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/campaign"
)

// Create 创建广告组
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.StatusUpdateRequest) ([]uint64, error) {
	var resp campaign.StatusUpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.CampaignIds, nil
}
