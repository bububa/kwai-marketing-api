package campaign

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/campaign"
	"github.com/bububa/kwai-marketing-api/model/v2.2/unit"
)

// 更新广告组预算
func BudgetUpdate(clt *core.SDKClient, accessToken string, req *unit.BudgetUpdateRequest) (uint64,error) {
	var resp campaign.BudgetUpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CampaignId, nil
}
