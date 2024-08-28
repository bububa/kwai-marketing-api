package campaign

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/campaign"
)

// Create 创建广告计划
// 【注】创建搜索广告计划。每个广告主账号下最多可允许创建500个计划
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) {
	var resp campaign.CreateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CampaignID, nil
}
