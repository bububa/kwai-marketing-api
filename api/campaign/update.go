package campaign

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/campaign"
)

// Update 修改广告计划
func Update(clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (int64, error) {
	var resp campaign.UpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CampaignID, nil
}
