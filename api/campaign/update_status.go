package campaign

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/campaign"
)

// UpdateStatus 修改广告计划状态
func UpdateStatus(clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) ([]uint64, error) {
	var resp campaign.UpdateStatusResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.CampaignIDs, nil
}
